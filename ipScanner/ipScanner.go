package ipScanner

import (
	"encoding/json"
	"io/ioutil"
	"io"
	"log"
	"net/http"
	"strconv"
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

//Receives the request that contain the client IP
func StartIPScan(response http.ResponseWriter, request *http.Request) {
	information := scanIP(ipClient(request))
	io.WriteString(response, information)
}

//The client IP will be his IP or the one that he writes in the url
func ipClient(request *http.Request) string {
	IP := request.URL.Query().Get("for")
	if IP == "" {
		IP = request.RemoteAddr
	}
	return IP
}

//Gets IP information using the freegeoip services
func scanIP(IP string) string {

	responseToIP, err := http.Get("http://freegeoip.net/json/" + IP)
	if err != nil {
		log.Fatal(err)
	}

	ipScanData, err := ioutil.ReadAll(responseToIP.Body)
	responseToIP.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	ipInfo := processIPScanInformation(ipScanData)

	return ipInfo
}

//Stores the IP information in a string  
func processIPScanInformation(ipScanData []byte) string {

	ipInformation := iPInformation{}
	json.Unmarshal(ipScanData, &ipInformation)

	Latitude := strconv.FormatFloat(ipInformation.Latitude, 'f', 6, 64)
	Longitude := strconv.FormatFloat(ipInformation.Longitude, 'f', 6, 64)
	ipInformation.URL = "https://www.google.com.au/maps/@" + Latitude + "," + Longitude + "z"

	ipInfo := "Host: "+ipInformation.IP+"\n"
	ipInfo += "Country: "+ipInformation.CountryName+"\nRegion: "+ipInformation.RegionName+"\n"
	ipInfo += "City: "+ipInformation.City+"\nZip Code: "+ipInformation.ZipCode+"\n"
	ipInfo += "Time Zone: "+ipInformation.TimeZone+"\nLink to google maps with geolocation: "+ipInformation.URL

	return ipInfo
}


