package utils

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

// GenerateNewAccessToken ... Generate new JWT token
// @Summary Generate new JWT token
// @Description Generates a new JWT token
// @Description and returns the token or error.
// @Tags Utils
func GenerateNewAccessToken(payload interface{}) (string, error) {
	secret := os.Getenv("JWT_SECRET_KEY")

	minutesCount, _ := strconv.Atoi(os.Getenv("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"))

	claims := jwt.MapClaims{}
	now := time.Now().UTC()

	claims["sub"] = payload
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
