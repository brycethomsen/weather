package main

import (
	"fmt"
	"io/ioutil"

	"github.com/brycethomsen/weather/pkg/nominatim"
)

func main() {
	city := "minneapolis"
	api, _ := ioutil.ReadFile("secret")

	lat, lon := nominatim.Geocode(city)

	// request := darksky.ForecastRequest{
	// 	Latitude:  lat,
	// 	Longitude: lon,
	// 	},
	// }
	fmt.Printf("%s%s%s", lat, lon, api)

	// forecast, err := darksky.Forecast(request)
}
