package response

import (
	"time"
)

type UserResponse struct {
	Id          uint      `json:"id"`
	Email       string    `json:"email"`
	DateOfBirth time.Time `json:"dateOfBirth"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
