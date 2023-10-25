package model

import (
	"time"
)

type User struct {
	Id                   uint                   `gorm:"primary_key;auto_increment"`
	Email                string                 `gorm:"type:varchar(100);unique;not null"`
	Password             string                 `gorm:"type:varchar(200);not null"`
	DateOfBirth          time.Time              `gorm:"type:date;not null"`
	WeatherSearchHistory []WeatherSearchHistory `gorm:"foreignKey:UserID;references:Id"`
}
