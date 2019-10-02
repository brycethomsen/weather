package nominatim

// ForecastResponse is the response containing all requested properties
type GetLatLon struct {
	Latitude  float64 `json:"lat,omitempty"`
	Longitude float64 `json:"lon,omitempty"`
}
