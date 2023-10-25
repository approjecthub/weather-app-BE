package helper

import (
	"net/http"
	"os"
	"time"
	"weather-app-BE/data/response"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func GenerateHash(data string) (string, error) {
	hashedData, err := bcrypt.GenerateFromPassword([]byte(data), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Msg(err.Error())
		return "", err
	}
	return string(hashedData), nil
}

func CompareHash(hashedData string, incomingData string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedData), []byte(incomingData))
	return err == nil
}

func ConverToValidDate(date string) (time.Time, error) {
	layout := "2006-01-02"

	// Attempt to parse the input string with the specified layout
	paredDate, err := time.Parse(layout, date)

	// If there's no error, it's a valid date in the specified format
	if err != nil {
		log.Error().Msg(err.Error())
		return time.Time{}, err
	}
	return paredDate, nil
}

func GenerateJWTToken(userId uint) (string, error) {
	secretKey := []byte(os.Getenv("SECRET_KEY"))
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Error().Msg(err.Error())
		return "", err
	}

	return tokenString, nil
}

func SendErrorResponse(err error, ctx *gin.Context) {
	log.Error().Msg(err.Error())
	errRes := response.Response{
		Data:  nil,
		Error: err.Error(),
	}
	ctx.JSON(http.StatusBadRequest, errRes)

}
