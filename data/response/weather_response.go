package response

import "time"

type CityResponse struct {
	Name    string  `json:"name"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
	Country string  `json:"country"`
	State   string  `json:"state"`
}

type WeatherSearchObject struct {
	Id          uint      `json:"id,omitempty"`
	Place       string    `validate:"required" json:"place,omitempty"`
	Lat         float64   `validate:"required" json:"lat,omitempty"`
	Lon         float64   `validate:"required" json:"lon,omitempty"`
	Country     string    `validate:"required" json:"country,omitempty"`
	Description string    `validate:"required" json:"description,omitempty"`
	Temparature float64   `validate:"required" json:"temparature,omitempty"`
	Wind        float64   `validate:"required" json:"wind,omitempty"`
	Humidity    float64   `validate:"required" json:"humidity,omitempty"`
	Pressure    float64   `validate:"required" json:"pressure,omitempty"`
	UserID      uint      `json:"userID,omitempty"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
}
