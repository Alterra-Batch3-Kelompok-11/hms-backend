package routes

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"hms-backend/configs"
	"hms-backend/controllers/authController"
	"hms-backend/repositories/doctorRepository"
	"hms-backend/repositories/nurseRepository"
	"hms-backend/repositories/userRepository"
	"hms-backend/usecases/authUseCase"
)

func New(db *gorm.DB) *echo.Echo {
	e := echo.New()
	configs.InitConfig()

	// Repositories
	usrRepo := userRepository.New(db)
	dtrRepo := doctorRepository.New(db)
	nrsRepo := nurseRepository.New(db)

	// Use Cases
	authUc := authUseCase.New(usrRepo, dtrRepo, nrsRepo)

	// Controllers
	authCtrl := authController.New(authUc)

	// Middlewares
	//jwt := middleware.JWT([]byte(configs.Cfg.JwtKey))
	//roleMdlwr := middlewares.RoleMiddleware

	// V1
	v1 := e.Group("/v1")
	v1.POST("/login", authCtrl.Login)
	v1.POST("/signup", authCtrl.SignUp)

	//e.GET("/generate-hash-password/:password", controllers.GenerateHashPassword)

	return e
}
