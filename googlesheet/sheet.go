package googlesheet

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/ledongthuc/hemnetparser"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

// Scope and Status of sheet line processing
const (
	ScopeURL   = "https://www.googleapis.com/auth/spreadsheets"
	StatusSkip = "Skipped"
)

// Default data of configuration
const (
	DefaultSheetName = "Sheet1"
	DefaultColumns   = "area_detail.postal_city,housing_form,price,selling_price,development_price,development_price_percent,number_of_rooms,living_area,,borattavgift,driftkostnad,construction_year,street_name"
)

// Configuration is config to process GoogleSheet updating
type Configuration struct {
	GoogleCredentialPath string
	SpreadsheetID        string
	NoSkipHeader         int
	SheetName            string
	Columns              string
}

// Validate check invalid in configuration
func (c Configuration) Validate() (bool, error) {
	if c.GoogleCredentialPath == "" {
		return false, fmt.Errorf("Missing --credential=<value>. Credential file path from https://developers.google.com/sheets/api/quickstart/go")
	}
	if c.SpreadsheetID == "" {
		return false, fmt.Errorf("Missing --sheet-id=<value>. https://developers.google.com/sheets/api/guides/concepts#spreadsheet_id")
	}
	return true, nil
}

// Default set default data of configuration if it doesn't set yet
func (c *Configuration) Default() {
	if c == nil {
		return
	}
	if c.SheetName == "" {
		c.SheetName = DefaultSheetName
	}
	if c.Columns == "" {
		c.Columns = DefaultColumns
	}
}

// Process starts to parse data from hemnet and update to Excel
func Process(configuration Configuration) ([]byte, error) {
	configuration.Default()
	b, err := ioutil.ReadFile(configuration.GoogleCredentialPath)
	if err != nil {
		return nil, fmt.Errorf("Unable to read client secret file: %v", configuration.GoogleCredentialPath)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, ScopeURL)
	if err != nil {
		return nil, fmt.Errorf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	srv, err := sheets.New(client)
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve Sheets client: %v", err)
	}

	// Prints the names and majors of students in a sample spreadsheet:
	// https://docs.google.com/spreadsheets/d/1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms/edit
	readRange := configuration.SheetName
	resp, err := srv.Spreadsheets.Values.Get(configuration.SpreadsheetID, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
		return nil, fmt.Errorf("Unable to retrieve data from sheet: %v", err)
	}

	if len(resp.Values) == 0 {
		return []byte("Success"), nil
	}

	var updatedValues [][]interface{}
	for index, row := range resp.Values {
		if index <= configuration.NoSkipHeader-1 {
			updatedValues = append(updatedValues, row)
			continue
		}
		updatedRow, err := processRow(row, strings.Split(configuration.Columns, ","))
		if err != nil {
			return []byte("Fail in index " + strconv.Itoa(index)), err
		}
		updatedValues = append(updatedValues, updatedRow)
	}

	_, err = srv.Spreadsheets.Values.Update(configuration.SpreadsheetID, readRange, &sheets.ValueRange{
		Values: updatedValues,
	}).ValueInputOption("USER_ENTERED").Do()
	if err != nil {
		return nil, err
	}

	return []byte("Success"), nil
}

func processRow(row []interface{}, columns []string) ([]interface{}, error) {
	var url string
	if len(row) > 0 {
		url = fmt.Sprintf("%v", row[0])
	}

	var currentStatus string
	if len(row) > 1 {
		currentStatus = fmt.Sprintf("%v", row[1])
	}
	if currentStatus == string(hemnetparser.StatusDeactived) ||
		currentStatus == string(hemnetparser.StatusSold) ||
		currentStatus == StatusSkip {
		return row, nil
	}

	output, err := hemnetparser.Parse(url)
	if err != nil {
		return row, err
	}

	response := []interface{}{
		url,
		output.Status,
	}
	for index, column := range columns {
		var defaultValue interface{}
		if len(row) >= index+1+2 {
			defaultValue = row[index+2]
		}
		if column == "" {
			response = append(response, defaultValue)
			continue
		}
		response = append(response, getDataFromColumn(strings.TrimSpace(column), output, defaultValue))
	}

	return response, nil
}

func getDataFromColumn(columnName string, hemnetData hemnetparser.Output, defaultValue interface{}) interface{} {
	switch columnName {
	case "url":
		return hemnetData.URL
	case "street_name":
		if hemnetData.StreetName == nil {
			return defaultValue
		}
		return *hemnetData.StreetName
	case "area":
		if hemnetData.Area == nil {
			return defaultValue
		}
		return *hemnetData.Area
	case "area_detail.postal_city":
		if hemnetData.AreaDetail.PostalCity == nil {
			return defaultValue
		}
		return *hemnetData.AreaDetail.PostalCity
	case "area_detail.municipality":
		if hemnetData.AreaDetail.Municipality == nil {
			return defaultValue
		}
		return *hemnetData.AreaDetail.Municipality
	case "area_detail.county":
		if hemnetData.AreaDetail.County == nil {
			return defaultValue
		}
		return *hemnetData.AreaDetail.County
	case "area_detail.country":
		if hemnetData.AreaDetail.Country == nil {
			return defaultValue
		}
		return *hemnetData.AreaDetail.Country
	case "construction_year":
		if hemnetData.ConstructionYear == nil {
			return defaultValue
		}
		return *hemnetData.ConstructionYear
	case "number_of_rooms":
		if hemnetData.NumberOfRooms == nil {
			return defaultValue
		}
		return *hemnetData.NumberOfRooms
	case "living_area":
		if hemnetData.NumberOfRooms == nil {
			return defaultValue
		}
		return hemnetData.LivingArea
	case "borattavgift":
		if hemnetData.Borattavgift == nil {
			return defaultValue
		}
		return *hemnetData.Borattavgift
	case "driftkostnad":
		if hemnetData.Driftkostnad == nil {
			return defaultValue
		}
		return *hemnetData.Driftkostnad
	case "price":
		if hemnetData.Price == nil {
			return defaultValue
		}
		return *hemnetData.Price
	case "price_per_m_2":
		if hemnetData.PricePerM2 == nil {
			return defaultValue
		}
		return *hemnetData.PricePerM2
	case "selling_price":
		if hemnetData.SellingPrice == nil {
			return defaultValue
		}
		return *hemnetData.SellingPrice
	case "development_price":
		if hemnetData.DevelopmentPrice == nil {
			return defaultValue
		}
		return *hemnetData.DevelopmentPrice
	case "development_price_percent":
		if hemnetData.DevelopmentPricePercent == nil {
			return defaultValue
		}
		return *hemnetData.DevelopmentPricePercent
	case "housing_form":
		if hemnetData.HousingForm == nil {
			return defaultValue
		}
		return *hemnetData.HousingForm
	case "tenure":
		if hemnetData.Tenure == nil {
			return defaultValue
		}
		return *hemnetData.Tenure
	}
	return ""
}

// getClient retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// getTokenFromWeb request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// tokenFromFile retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// saveToken saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}
