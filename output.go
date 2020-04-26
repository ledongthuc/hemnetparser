package hemnetparser

import (
	"encoding/json"
	"encoding/xml"
)

// Output is parsed output from library
type Output struct {
	URL              string     `json:"url,omitempty"`
	StreetName       *string    `json:"street_name,omitempty"`
	Area             *string    `json:"area,omitempty"`
	AreaDetail       AreaDetail `json:"area_detail"`
	ConstructionYear *string    `json:"construction_year,omitempty"`
	NumberOfRooms    *float64   `json:"number_of_rooms,omitempty"`
	LivingArea       *float64   `json:"living_area,omitempty"`
	Borattavgift     *float64   `json:"borattavgift,omitempty"`
	Driftkostnad     *float64   `json:"driftkostnad,omitempty"`
	Price            *float64   `json:"price,omitempty"`
	PricePerM2       *float64   `json:"price_per_m_2,omitempty"`
	HousingForm      *string    `json:"housing_form",omitempty`
	Tenure           *string    `json:"tenure,omitempty"`
}

// AreaDetail define  area information such as city, county or country
type AreaDetail struct {
	PostalCity   *string `json:"postal_city,omitempty"`
	Municipality *string `json:"municipality,omitempty"`
	County       *string `json:"county,omitempty"`
	Country      *string `json:"country,omitempty"`
}

// ToJSON convert output object to JSON
func (o Output) ToJSON() ([]byte, error) {
	return json.MarshalIndent(o, "", " ")
}

// ToXML convert output object to JSON
func (o Output) ToXML() ([]byte, error) {
	return xml.MarshalIndent(o, "", " ")
}
