package hemnetparser

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Parse parses information from URL to output object
func Parse(rawURL string) (Output, error) {
	res, err := http.Get(rawURL)
	if err != nil {
		return Output{}, fmt.Errorf("Error when get data from url %v: %v", rawURL, err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		return Output{}, fmt.Errorf("Got error code when loading url: %d, %s", res.StatusCode, res.Status)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Output{}, fmt.Errorf("Got read data from website: %v, %s", rawURL, err)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(data)))
	if err != nil {
		return Output{}, fmt.Errorf("Got read data from website: %v, %s", rawURL, err)
	}

	re := regexp.MustCompile("\"construction_year\":\"(\\w+)\"")
	matches := re.FindStringSubmatch(string(data))
	constructionYear := ""
	if len(matches) > 0 {
		constructionYear = matches[len(matches)-1]
	}

	// Check sold property
	if soldButton := doc.Find(".removed-listing__button"); soldButton != nil {
		if v, exists := soldButton.Attr("href"); exists {
			if !strings.HasPrefix(v, "http") {
				if resURL, err := res.Location(); err == nil && resURL != nil {
					v = fmt.Sprintf("%s://%s/%s", resURL.Scheme, resURL.Host, v)
				}
			}
			return Parse(v)
		}
	}

	// Find the review items
	output := Output{
		URL:        rawURL,
		StreetName: StringValueOrNil(doc.Find(".property-address__street").Text()),
		Area:       StringValueOrNil(doc.Find(".property-address__area").Text()),
		AreaDetail: AreaDetail{
			PostalCity:   StringValueOrNil(parseJson("postal_city", data)),
			Municipality: StringValueOrNil(parseJson("municipality", data)),
			County:       StringValueOrNil(parseJson("county", data)),
			Country:      StringValueOrNil(parseJson("country", data)),
		},
		ConstructionYear: StringValueOrNil(constructionYear),
		HousingForm:      StringValueOrNil(parseJson("housing_form", data)),
		Tenure:           StringValueOrNil(parseJson("tenure", data)),
		Status:           parseStatus(data),
	}
	if output.StreetName == nil || *output.StreetName == "" {
		output.StreetName = StringValueOrNil(parseJson("streetAddress", data))
	}
	if v, err := strconv.ParseFloat(parseJson("rooms", data), 32); err == nil {
		output.NumberOfRooms = &v
	}
	if v, err := strconv.ParseFloat(parseJson("living_area", data), 32); err == nil {
		output.LivingArea = &v
	}
	if v, err := strconv.ParseFloat(parseJson("borattavgift", data), 64); err == nil {
		output.Borattavgift = &v
	}
	if v, err := strconv.ParseFloat(parseJson("driftkostnad", data), 64); err == nil {
		output.Driftkostnad = &v
	}
	if v, err := strconv.ParseFloat(parseJson("price", data), 64); err == nil {
		output.Price = &v
	}
	output.SellingPrice = parseSellingPrice(data, doc, output.Status)
	if v, err := strconv.ParseFloat(parseJson("price_per_m2", data), 64); err == nil {
		output.PricePerM2 = &v
	}
	output.CalculateDevelopment()
	return output, nil
}

// parseJson value from json key in a document
func parseJson(key string, documentB []byte) string {
	document := string(documentB)
	re := regexp.MustCompile("\"" + key + "\"\\s*:\\s*")
	matches := re.FindStringIndex(document)
	if len(matches) != 2 {
		return ""
	}

	var value strings.Builder
	var lastChar byte
	inQuotes := false
	for i := matches[1]; i < len(document); i++ {
		lastChar = document[i]
		if !inQuotes && (document[i] == ',' || document[i] == '}') {
			// "key":"value", ...
			// "key":"value"} ...
			break
		}
		if document[i] == '"' {
			if lastChar != '\\' {
				inQuotes = !inQuotes
				continue
			}
		}
		value.WriteByte(document[i])
	}
	return value.String()
}

func parseStatus(documentB []byte) Status {
	if strings.Contains(string(documentB), `"status":"deactivated"`) ||
		strings.Contains(string(documentB), `"status":"deactivated_before_open_house_day"`) {
		return StatusDeactived
	}
	if strings.Contains(string(documentB), `&quot;status&quot;:&quot;sold&quot;`) {
		return StatusSold
	}
	return StatusProcessing
}

func parseSellingPrice(documentB []byte, doc *goquery.Document, status Status) *float64 {
	if v, err := strconv.ParseFloat(parseJson("selling_price", documentB), 64); err == nil {
		return &v
	}
	if status == StatusDeactived {
		rawSellingPrice := doc.Find(".removed-listing__price").Text()
		rawSellingPrice = stripSpaces(strings.ReplaceAll(rawSellingPrice, "kr", ""))
		if v, err := strconv.ParseFloat(rawSellingPrice, 64); err == nil {
			return &v
		}
	}
	return nil
}
