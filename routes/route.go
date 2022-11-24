package routes

import (
	"hms-backend/configs"
	"hms-backend/controllers/authController"
	"hms-backend/controllers/roleController"
	controllers "hms-backend/controllers/userController"
	"hms-backend/repositories/doctorRepository"
	"hms-backend/repositories/nurseRepository"
	"hms-backend/repositories/roleRepository"
	"hms-backend/repositories/userRepository"
	"hms-backend/usecases/authUseCase"
	"hms-backend/usecases/roleUseCase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func New(db *gorm.DB, echoSwagger echo.HandlerFunc) *echo.Echo {
	e := echo.New()
	configs.InitConfig()

	// Repositories
	usrRepo := userRepository.New(db)
	dtrRepo := doctorRepository.New(db)
	nrsRepo := nurseRepository.New(db)
	rlRepo := roleRepository.New(db)

	// Use Cases
	authUc := authUseCase.New(usrRepo, dtrRepo, nrsRepo)
	rlUc := roleUseCase.New(rlRepo)

	// Controllers
	authCtrl := authController.New(authUc)
	rlCtrl := roleController.New(rlUc)

	// Middlewares
	//jwt := middleware.JWT([]byte(configs.Cfg.JwtKey))
	//roleMdlwr := middlewares.RoleMiddleware

	e.GET("/swagger/*", echoSwagger)

	// V1
	v1 := e.Group("/v1")
	v1.POST("/login", authCtrl.Login)
	v1.POST("/signup", authCtrl.SignUp)

	role := v1.Group("/roles")
	role.GET("", rlCtrl.GetAll)
	role.GET("/:id", rlCtrl.GetById)

	// CRUD Patients

	v1.POST("/createpatient", controllers.PatientsCreate)
	v1.GET("/indexpatient", controllers.PatientsIndex)
	v1.GET("/indexpatient/:id", controllers.PatientShow)
	v1.PUT("/updatepatient/:id", controllers.PatientsUpdate)
	v1.DELETE("/deletepatient/:id", controllers.PatientsDelete)

	return e
}
