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

const (
	ScopeURL        = "https://www.googleapis.com/auth/spreadsheets"
	StatusProcessed = "Processed"
	StatusSkip      = "Skipped"
)

type Configuration struct {
	GoogleCredentialPath string
	SpreadsheetID        string
	NoSkipHeader         int
	SheetName            string
	Columns              string
}

func (c Configuration) Validate() (bool, error) {
	if c.GoogleCredentialPath == "" {
		return false, fmt.Errorf("Missing --credential=<value>. Credential file path from https://developers.google.com/sheets/api/quickstart/go")
	}
	if c.SpreadsheetID == "" {
		return false, fmt.Errorf("Missing --sheet-id=<value>. https://developers.google.com/sheets/api/guides/concepts#spreadsheet_id")
	}
	return true, nil
}

func (c *Configuration) Default() {
	if c == nil {
		return
	}
	if c.SheetName == "" {
		c.SheetName = "Sheet1"
	}
	if c.Columns == "" {
		c.Columns = "area_detail.postal_city,housing_form,price,,,,number_of_rooms,living_area,,borattavgift,driftkostnad,construction_year,street_name"
	}
}

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

	var status string
	if len(row) > 1 {
		status = fmt.Sprintf("%v", row[1])
	}
	if status == StatusProcessed || status == StatusSkip {
		return row, nil
	}

	output, err := hemnetparser.Parse(url)
	if err != nil {
		return row, err
	}

	status = StatusProcessed
	response := []interface{}{
		url,
		status,
	}
	for _, column := range columns {
		response = append(response, getDataFromColumn(strings.TrimSpace(column), output))
	}

	return response, nil
}

func getDataFromColumn(columnName string, hemnetData hemnetparser.Output) interface{} {
	switch columnName {
	case "url":
		return hemnetData.URL
	case "street_name":
		return hemnetData.StreetName
	case "area":
		return hemnetData.Area
	case "area_detail.postal_city":
		return hemnetData.AreaDetail.PostalCity
	case "area_detail.municipality":
		return hemnetData.AreaDetail.Municipality
	case "area_detail.county":
		return hemnetData.AreaDetail.County
	case "area_detail.country":
		return hemnetData.AreaDetail.Country
	case "construction_year":
		return hemnetData.ConstructionYear
	case "number_of_rooms":
		return hemnetData.NumberOfRooms
	case "living_area":
		return hemnetData.LivingArea
	case "borattavgift":
		return hemnetData.Borattavgift
	case "driftkostnad":
		return hemnetData.Driftkostnad
	case "price":
		return hemnetData.Price
	case "price_per_m_2":
		return hemnetData.PricePerM2
	case "housing_form":
		return hemnetData.HousingForm
	case "tenure":
		return hemnetData.Tenure
	}
	return ""
}

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
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

// Retrieves a token from a local file.
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

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}
