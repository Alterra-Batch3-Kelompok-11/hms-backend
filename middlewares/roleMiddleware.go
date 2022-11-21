package middlewares

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"hms-backend/configs"
	"net/http"
	"strings"
)

func RoleMiddleware(next echo.HandlerFunc, allowedRoles [...]uint) echo.HandlerFunc {
	return func(c echo.Context) error {
		headerToken := c.Request().Header.Get("Authorization")
		token := strings.Split(headerToken, " ")[1]
		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
			return []byte(configs.Cfg.JwtKey), nil
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Unable to parse token")
		}

		isAllowed := true
		if len(allowedRoles) > 0 {
			isAllowed = false
			for _, allowedRole := range allowedRoles {
				if allowedRole == claims["roleId"].(uint) {
					isAllowed = true
				}
			}
		}
		if !isAllowed {
			return echo.NewHTTPError(http.StatusUnauthorized, "Not authorized")
		}

		return next(c)
	}
}
