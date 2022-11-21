package routes

import (
	"hms-backend/configs"
	controllers "hms-backend/controllers/userController"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *echo.Echo {
	e := echo.New()
	configs.InitConfig()

	e.GET("/generate-hash-password/:password", controllers.GenerateHashPassword)

	return e
}
