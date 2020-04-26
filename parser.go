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
func Parse(url string) (Output, error) {
	res, err := http.Get(url)
	if err != nil {
		return Output{}, fmt.Errorf("Error when get data from url %v: %v", url, err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		return Output{}, fmt.Errorf("Got error code when loading url: %d, %s", res.StatusCode, res.Status)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Output{}, fmt.Errorf("Got read data from website: %v, %s", url, err)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(data)))
	if err != nil {
		return Output{}, fmt.Errorf("Got read data from website: %v, %s", url, err)
	}

	re := regexp.MustCompile("\"construction_year\":\"(\\w+)\"")
	matches := re.FindStringSubmatch(string(data))
	constructionYear := ""
	if len(matches) > 0 {
		constructionYear = matches[len(matches)-1]
	}

	// Find the review items
	output := Output{
		URL:        url,
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
	}
	if v, err := strconv.ParseFloat(parseJson("rooms", data), 32); err != nil {
		output.NumberOfRooms = &v
	}
	if v, err := strconv.ParseFloat(parseJson("living_area", data), 32); err != nil {
		output.LivingArea = &v
	}
	if v, err := strconv.ParseFloat(parseJson("borattavgift", data), 64); err != nil {
		output.Borattavgift = &v
	}
	if v, err := strconv.ParseFloat(parseJson("driftkostnad", data), 64); err != nil {
		output.Driftkostnad = &v
	}
	if v, err := strconv.ParseFloat(parseJson("price", data), 64); err != nil {
		output.Price = &v
	}
	if v, err := strconv.ParseFloat(parseJson("price_per_m2", data), 64); err != nil {
		output.PricePerM2 = &v
	}
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
		if !inQuotes && document[i] == ',' {
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
