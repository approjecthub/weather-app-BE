package model

import "time"

type WeatherSearchHistory struct {
	Id          uint      `gorm:"primary_key;auto_increment"`
	Place       string    `gorm:"type:varchar(50);not null"`
	Lat         float64   `gorm:"type:decimal(20,10);not null"`
	Lon         float64   `gorm:"type:decimal(20,10);not null"`
	Country     string    `gorm:"type:varchar(50);not null"`
	Description string    `gorm:"type:string"`
	Temparature float64   `gorm:"type:decimal(20,10);not null"`
	Wind        float64   `gorm:"type:decimal(20,10);not null"`
	Humidity    float64   `gorm:"type:decimal(20,10);not null"`
	Pressure    float64   `gorm:"type:decimal(20,10);not null"`
	UserID      uint      `gorm:"not null"`
	CreatedAt   time.Time `gorm:"type:datetime;not null"`
}
