package models

type GeocodeResponse struct {
	Results []struct {
		Geometry struct {
			Location struct {
				Latitude  float64 `json:"lat"`
				Longitude float64 `json:"lng"`
			} `json:"location"`
		} `json:"geometry"`
	} `json:"results"`
}

type AQIResponse struct {
	Indexes []struct {
		AQI int `json:"aqi"`
	} `json:"indexes"`
}
