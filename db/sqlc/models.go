// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import ()

type EcdeCase struct {
	ID                      int64  `json:"id"`
	DateRep                 string `json:"date_rep"`
	Day                     int32  `json:"day"`
	Month                   int32  `json:"month"`
	Year                    int32  `json:"year"`
	Cases                   int64  `json:"cases"`
	Deaths                  int64  `json:"deaths"`
	CountriesAndTerritories string `json:"countries_and_territories"`
	GeoID                   string `json:"geo_id"`
	CountryTerritoryCode    string `json:"country_territory_code"`
	ContinentExp            string `json:"continent_exp"`
	LoadDate                string `json:"load_date"`
	IsoCountry              string `json:"iso_country"`
}
