package service

type InfoIP struct {
	Ip 				string	`json:"ip"`
	Ip_type 		string	`json:"type"`
	Continent_code 	string	`json:"continent_code"`
	Continent_name 	string	`json:"continent_name"`
	Country_code 	string	`json:"country_code"`
	Country_name 	string	`json:"country_name"`
	Region_code 	string 	`json:"region_code"`
	Region_name 	string	`json:"region_name"`
	City 			string	`json:"city"`
	Zip 			string	`json:"zip"`
	Latitude 		float64	`json:"latitude"`
	Longitude 		float64	`json:"longitude"`
}
