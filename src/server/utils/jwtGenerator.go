package utils

import (
	"strconv"
	"time"

	"github.com/HRemonen/kanban-board/config"
	"github.com/golang-jwt/jwt"
)

func GenerateNewAccessToken() (string, error) {
	secret := config.Config("JWT_SECRET_KEY")

	minutesCount, _ := strconv.Atoi(config.Config("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"))

	claims := jwt.MapClaims{}

	claims["exp"] = time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return t, nil
}
