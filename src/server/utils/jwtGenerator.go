package utils

import (
	"strconv"
	"time"

	"github.com/HRemonen/kanban-board/config"
	"github.com/golang-jwt/jwt"
)

func GenerateNewAccessToken(payload interface{}) (string, error) {
	secret := config.Config("JWT_SECRET_KEY")

	minutesCount, _ := strconv.Atoi(config.Config("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"))

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
