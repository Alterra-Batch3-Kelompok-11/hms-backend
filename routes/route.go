package routes

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"hms-backend/configs"
)

func New(db *gorm.DB) *echo.Echo {
	e := echo.New()
	configs.InitConfig()

	return e
}
