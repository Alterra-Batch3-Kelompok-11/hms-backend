package routes

import (
	"github.com/labstack/echo/v4/middleware"
	"hms-backend/configs"
	"hms-backend/controllers/authController"
	"hms-backend/controllers/roleController"
	"hms-backend/controllers/specialityController"
	controllers "hms-backend/controllers/userController"
	"hms-backend/middlewares"
	"hms-backend/repositories/doctorRepository"
	"hms-backend/repositories/nurseRepository"
	"hms-backend/repositories/roleRepository"
	"hms-backend/repositories/specialityRepository"
	"hms-backend/repositories/userRepository"
	"hms-backend/usecases/authUseCase"
	"hms-backend/usecases/roleUseCase"
	"hms-backend/usecases/specialityUseCase"

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
	spcRepo := specialityRepository.New(db)

	// Use Cases
	authUc := authUseCase.New(usrRepo, dtrRepo, nrsRepo)
	rlUc := roleUseCase.New(rlRepo)
	spcUc := specialityUseCase.New(spcRepo)

	// Controllers
	authCtrl := authController.New(authUc)
	rlCtrl := roleController.New(rlUc)
	spcCtrl := specialityController.New(spcUc)

	// Middlewares
	jwt := middleware.JWT([]byte(configs.Cfg.JwtKey))
	admMdlwr := middlewares.RoleAdminMiddleware
	//dctrMdlwr := middlewares.RoleDoctorMiddleware
	//nrsMdlwr := middlewares.RoleNurseMiddleware

	e.GET("/swagger/*", echoSwagger)

	// V1
	v1 := e.Group("/v1")
	v1.POST("/login", authCtrl.Login)
	v1.POST("/signup", authCtrl.SignUp)

	// Roles
	role := v1.Group("/roles")
	role.GET("", rlCtrl.GetAll)
	role.GET("/:id", rlCtrl.GetById)

	// Specialities
	specialty := v1.Group("/specialities")
	specialty.GET("", spcCtrl.GetAll)
	specialty.GET("/:id", spcCtrl.GetById)
	specialty.POST("", spcCtrl.Create, jwt, admMdlwr)
	specialty.PUT("/:id", spcCtrl.Update, jwt, admMdlwr)
	specialty.DELETE("/:id", spcCtrl.Delete, jwt, admMdlwr)

	// CRUD Patients

	v1.POST("/createpatient", controllers.PatientsCreate)
	v1.GET("/indexpatient", controllers.PatientsIndex)
	v1.GET("/indexpatient/:id", controllers.PatientShow)
	v1.PUT("/updatepatient/:id", controllers.PatientsUpdate)
	v1.DELETE("/deletepatient/:id", controllers.PatientsDelete)

	return e
}
