package hemnetparser

import (
	"encoding/json"
	"encoding/xml"
)

// Output is parsed output from library
type Output struct {
	URL              string     `json:"url"`
	StreetName       string     `json:"street_name"`
	Area             string     `json:"area"`
	AreaDetail       AreaDetail `json:"area_detail"`
	ConstructionYear string     `json:"construction_year"`
	NumberOfRooms    float64    `json:"number_of_rooms"`
	LivingArea       float64    `json:"living_area"`
	Borattavgift     float64    `json:"borattavgift"`
	Driftkostnad     float64    `json:"driftkostnad"`
	Price            float64    `json:"price"`
	PricePerM2       float64    `json:"price_per_m_2"`
	HousingForm      string     `json:"housing_form"`
	Tenure           string     `json:"tenure"`
}

// AreaDetail define  area information such as city, county or country
type AreaDetail struct {
	PostalCity   string `json:"postal_city"`
	Municipality string `json:"municipality"`
	County       string `json:"county"`
	Country      string `json:"country"`
}

// ToJSON convert output object to JSON
func (o Output) ToJSON() ([]byte, error) {
	return json.MarshalIndent(o, "", " ")
}

// ToXML convert output object to JSON
func (o Output) ToXML() ([]byte, error) {
	return xml.MarshalIndent(o, "", " ")
}
