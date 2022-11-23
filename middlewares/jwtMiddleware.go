package middlewares

import (
	"github.com/golang-jwt/jwt/v4"
	"hms-backend/configs"
	"time"
)

func CreateToken(userId uint, username string, role uint) (string, error) {
	configs.InitConfig()

	claims := jwt.MapClaims{}
	claims["userId"] = userId
	claims["username"] = username
	claims["roleId"] = role
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(configs.Cfg.JwtKey))
}
