package nominatim

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// LatLon is latittude and longitude from nominatim OSM json response
type LatLon []struct {
	Latitude  string `json:"lat,omitempty"`
	Longitude string `json:"lon,omitempty"`
}

// Geocode will convert a cityname into lat/lon for darksky
func Geocode(city string) (string, string) {
	url := fmt.Sprintf("https://nominatim.openstreetmap.org/?format=json&city=%s&format=json&limit=1", city)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	geocode, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	var latlon LatLon
	json.Unmarshal([]byte(geocode), &latlon)
	if err != nil {
		log.Println(err)
	}
	lat := fmt.Sprintf("%s", latlon[0].Latitude)
	lon := fmt.Sprintf("%s", latlon[0].Longitude)
	return lat, lon
}
