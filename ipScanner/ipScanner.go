package ipScanner

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type iPInformation struct {
	IP          string
	CountryCode string `json:"country_code"`
	CountryName string `json:"country_name"`
	RegionCode  string `json:"region_code"`
	RegionName  string `json:"region_name"`
	City        string `json:"city"`
	ZipCode     string `json:"zip_code"`
	TimeZone    string `json:"time_zone"`
	Latitude    float64 
	Longitude   float64 
    URL         string 
}

func ScanIP (IP string) iPInformation {

	responseToIP, err := http.Get("http://freegeoip.net/json/"+IP)
	if err != nil {
		log.Fatal(err)
	}

	ipScanData, err := ioutil.ReadAll(responseToIP.Body)
	responseToIP.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	ipInformation := iPInformation{}
	json.Unmarshal(ipScanData, &ipInformation)

    //Latitude := (string) ipInformation.Latitude

    //ipInformation.URL ="https://www.google.com.au/maps/@"+ (string) ipInformation.Latitude+","+ (string) ipInformation.Latitude
	
	//url := ""+"https://www.google.com.au/maps/@"+iPInformation.Longitude+","+iPInformation.Longitude+"z"
	//url := "https://www.google.com.au/maps/@"+ipInformation.Latitude+","+ipInformation.Longitude+","
    return ipInformation;
}