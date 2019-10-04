package darksky

// ForecastRequest contains all available options for requesting a forecast
type ForecastRequest struct {
	Latitude  float64
	Longitude float64
	Time      int64
	Options   ForecastRequestOptions
}

// ForecastRequestOptions are optional and passed as query parameters
type ForecastRequestOptions struct {
	Exclude string
	Extend  string
	Lang    string
	Units   string
}

// ForecastResponse is the response containing all requested properties
type ForecastResponse struct {
	Latitude  float64    `json:"latitude,omitempty"`
	Longitude float64    `json:"longitude,omitempty"`
	Timezone  string     `json:"timezone,omitempty"`
	Currently *DataPoint `json:"currently,omitempty"`
	Minutely  *DataBlock `json:"minutely,omitempty"`
	Hourly    *DataBlock `json:"hourly,omitempty"`
	Daily     *DataBlock `json:"daily,omitempty"`
	Alerts    []*Alert   `json:"alerts,omitempty"`
	Flags     *Flags     `json:"flags,omitempty"`
}

// DataPoint contains various properties, each representing the average (unless otherwise specified) of a particular weather phenomenon occurring during a period of time.
type DataPoint struct {
	ApparentTemperature         float64 `json:"apparentTemperature,omitempty"`
	ApparentTemperatureHigh     float64 `json:"apparentTemperatureHigh,omitempty"`
	ApparentTemperatureHighTime int64   `json:"apparentTemperatureHighTime,omitempty"`
	ApparentTemperatureLow      float64 `json:"apparentTemperatureLow,omitempty"`
	ApparentTemperatureLowTime  int64   `json:"apparentTemperatureLowTime,omitempty"`
	ApparentTemperatureMax      float64 `json:"apparentTemperatureMax,omitempty"`
	ApparentTemperatureMaxTime  int64   `json:"apparentTemperatureMaxTime,omitempty"`
	ApparentTemperatureMin      float64 `json:"apparentTemperatureMin,omitempty"`
	ApparentTemperatureMinTime  int64   `json:"apparentTemperatureMinTime,omitempty"`
	CloudCover                  float64 `json:"cloudCover,omitempty"`
	DewPoint                    float64 `json:"dewPoint,omitempty"`
	Humidity                    float64 `json:"humidity,omitempty"`
	Icon                        string  `json:"icon,omitempty"`
	MoonPhase                   float64 `json:"moonPhase,omitempty"`
	NearestStormBearing         float64 `json:"nearestStormBearing,omitempty"`
	NearestStormDistance        float64 `json:"nearestStormDistance,omitempty"`
	Ozone                       float64 `json:"ozone,omitempty"`
	PrecipAccumulation          float64 `json:"precipAccumulation,omitempty"`
	PrecipIntensity             float64 `json:"precipIntensity,omitempty"`
	PrecipIntensityError        float64 `json:"precipIntensityError,omitempty"`
	PrecipIntensityMax          float64 `json:"precipIntensityMax,omitempty"`
	PrecipIntensityMaxTime      int64   `json:"precipIntensityMaxTime,omitempty"`
	PrecipProbability           float64 `json:"precipProbability,omitempty"`
	PrecipType                  string  `json:"precipType,omitempty"`
	Pressure                    float64 `json:"pressure,omitempty"`
	Summary                     string  `json:"summary,omitempty"`
	SunriseTime                 int64   `json:"sunriseTime,omitempty"`
	SunsetTime                  int64   `json:"sunsetTime,omitempty"`
	Temperature                 float64 `json:"temperature,omitempty"`
	TemperatureHigh             float64 `json:"temperatureHigh,omitempty"`
	TemperatureHighTime         int64   `json:"temperatureHighTime,omitempty"`
	TemperatureLow              float64 `json:"temperatureLow,omitempty"`
	TemperatureLowTime          int64   `json:"temperatureLowTime,omitempty"`
	TemperatureMax              float64 `json:"temperatureMax,omitempty"`
	TemperatureMaxTime          int64   `json:"temperatureMaxTime,omitempty"`
	TemperatureMin              float64 `json:"temperatureMin,omitempty"`
	TemperatureMinTime          int64   `json:"temperatureMinTime,omitempty"`
	Time                        int64   `json:"time,omitempty"`
	UvIndex                     int64   `json:"uvIndex,omitempty"`
	UvIndexTime                 int64   `json:"uvIndexTime,omitempty"`
	Visibility                  float64 `json:"visibility,omitempty"`
	WindBearing                 float64 `json:"windBearing,omitempty"`
	WindGust                    float64 `json:"windGust,omitempty"`
	WindGustTime                int64   `json:"windGustTime,omitempty"`
	WindSpeed                   float64 `json:"windSpeed,omitempty"`
}

// DarkSky API endpoint
var (
	BaseUrl = "https://api.darksky.net/forecast"
)

// Forecast request a forecast
func (d *darkSky) Forecast(request ForecastRequest) (ForecastResponse, error) {
	response := ForecastResponse{}

	requestUrl := d.buildRequestUrl(request)

	err := get(d.Client, requestUrl, &response)

	return response, err
}

func (d *darkSky) buildRequestUrl(request ForecastRequest) string {
	requestUrl := fmt.Sprintf("%s/%s/%g,%g", BaseUrl, d.ApiKey, request.Latitude, request.Longitude)

	if request.Time > 0 {
		requestUrl = requestUrl + fmt.Sprintf(",%d", request.Time)
	}

	queryString := request.Options.Encode()
	if queryString != "" {
		requestUrl = requestUrl + "?" + queryString
	}

	return requestUrl
}