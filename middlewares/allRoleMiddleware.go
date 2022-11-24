package middlewares

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"hms-backend/configs"
	"net/http"
	"strings"
)

func AllRoleMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
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

		allowedRoles := [3]uint{
			1, // Admin
			2, // Doctor
			3, // Nurse
		}

		isAllowed := false
		for _, allowedRole := range allowedRoles {
			if allowedRole == uint(claims["roleId"].(float64)) {
				isAllowed = true
			}
		}

		if !isAllowed {
			return echo.NewHTTPError(http.StatusUnauthorized, "Not authorized")
		}

		return next(c)
	}
}
