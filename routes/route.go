package routes

import (
	"hms-backend/configs"
	"hms-backend/controllers/authController"
	"hms-backend/controllers/doctorController"
	"hms-backend/controllers/doctorScheduleController"
	"hms-backend/controllers/nurseController"
	"hms-backend/controllers/outpatientSessionController"
	"hms-backend/controllers/patientController"
	"hms-backend/controllers/religionController"
	"hms-backend/controllers/roleController"
	"hms-backend/controllers/specialityController"
	"hms-backend/middlewares"
	"hms-backend/repositories/doctorRepository"
	"hms-backend/repositories/doctorScheduleRepository"
	"hms-backend/repositories/nurseRepository"
	"hms-backend/repositories/outpatientSessionRepository"
	"hms-backend/repositories/patientRepository"
	"hms-backend/repositories/religionRepository"
	"hms-backend/repositories/roleRepository"
	"hms-backend/repositories/specialityRepository"
	"hms-backend/repositories/userRepository"
	"hms-backend/usecases/authUseCase"
	"hms-backend/usecases/doctorScheduleUseCase"
	"hms-backend/usecases/doctorUseCase"
	"hms-backend/usecases/nurseUseCase"
	"hms-backend/usecases/outpatientSessionUseCase"
	"hms-backend/usecases/patientUseCase"
	"hms-backend/usecases/religionUseCase"
	"hms-backend/usecases/roleUseCase"
	"hms-backend/usecases/specialityUseCase"
	"net/http"

	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func New(db *gorm.DB, echoSwagger echo.HandlerFunc) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderAccessControlAllowOrigin, echo.HeaderContentType},
	}))

	configs.InitConfig()

	// Repositories
	usrRepo := userRepository.New(db)
	dtrRepo := doctorRepository.New(db)
	nrsRepo := nurseRepository.New(db)
	rlRepo := roleRepository.New(db)
	spcRepo := specialityRepository.New(db)
	patRepo := patientRepository.New(db)
	rlgRepo := religionRepository.New(db)
	dtrSchedRepo := doctorScheduleRepository.New(db)
	nurRepo := nurseRepository.New(db)
	outPatientSessionRepo := outpatientSessionRepository.New(db)

	// Use Cases
	authUc := authUseCase.New(usrRepo, dtrRepo, nrsRepo)
	rlUc := roleUseCase.New(rlRepo)
	spcUc := specialityUseCase.New(spcRepo)
	patUc := patientUseCase.New(patRepo)
	rlgUc := religionUseCase.New(rlgRepo)
	dtrUc := doctorUseCase.New(dtrRepo, usrRepo, spcRepo, dtrSchedRepo)
	dtrSchdUc := doctorScheduleUseCase.New(dtrRepo, dtrSchedRepo)
	nurUC := nurseUseCase.New(nurRepo)
	outPatientSessionUC := outpatientSessionUseCase.New(outPatientSessionRepo, usrRepo, dtrRepo, spcRepo, dtrSchedRepo, patRepo)

	// Controllers
	authCtrl := authController.New(authUc)
	rlCtrl := roleController.New(rlUc)
	spcCtrl := specialityController.New(spcUc)
	patCtrl := patientController.New(patUc)
	rlgCtrl := religionController.New(rlgUc)
	dtrCtrl := doctorController.New(dtrUc)
	dtrSchdCtrl := doctorScheduleController.New(dtrSchdUc)
	nurCtrl := nurseController.New(nurUC)
	outpatientSessionCtrl := outpatientSessionController.New(outPatientSessionUC)

	// Middlewares
	jwt := middleware.JWT([]byte(configs.Cfg.JwtKey))
	admMdlwr := middlewares.RoleAdminMiddleware
	dctrMdlwr := middlewares.RoleDoctorMiddleware
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

	// Doctors
	doctor := v1.Group("/doctors")
	doctor.GET("", dtrCtrl.GetAll)
	doctor.GET("/:id", dtrCtrl.GetById)
	doctor.GET("/speciality/:speciality_id", dtrCtrl.GetBySpecialityId)
	doctor.GET("/license_number/:license_number", dtrCtrl.GetByLicenseNumber)
	doctor.GET("/today", dtrCtrl.GetToday)
	doctor.POST("", dtrCtrl.Create, jwt, admMdlwr)
	doctor.PUT("/:id", dtrCtrl.Update, jwt, admMdlwr)
	doctor.DELETE("/:id", dtrCtrl.Delete, jwt, admMdlwr)

	// Doctor Schedules
	doctorSchedule := v1.Group("/doctor_schedules")
	doctorSchedule.GET("/doctor/:doctor_id", dtrSchdCtrl.GetByDoctorId)
	doctorSchedule.GET("/:id", dtrSchdCtrl.GetById)
	doctorSchedule.GET("/doctor/license_number/:license_number", dtrSchdCtrl.GetByLicenseNumber)
	doctorSchedule.POST("", dtrSchdCtrl.Create, jwt, admMdlwr)
	doctorSchedule.PUT("/:id", dtrSchdCtrl.Update, jwt, admMdlwr)
	doctorSchedule.DELETE("/:id", dtrSchdCtrl.Delete, jwt, admMdlwr)

	// Religions
	religion := v1.Group("/religions")
	religion.GET("", rlgCtrl.GetAll)
	religion.GET("/:id", rlgCtrl.GetById)

	// Patients
	patient := v1.Group("/patients")
	patient.GET("", patCtrl.GetAll)
	patient.GET("/:id", patCtrl.GetById)
	patient.POST("", patCtrl.Create, jwt, admMdlwr)
	patient.PUT("/:id", patCtrl.Update, jwt, admMdlwr)
	patient.DELETE("/:id", patCtrl.Delete, jwt, admMdlwr)

	// Nurses
	nurse := v1.Group("/nurses")
	nurse.GET("", nurCtrl.GetAll)
	nurse.GET("/:id", nurCtrl.GetById)
	nurse.GET("/license_number/:license_number", nurCtrl.GetByLicenseNumber)
	nurse.POST("", nurCtrl.Create, jwt, admMdlwr)
	nurse.PUT("/:id", nurCtrl.Update, jwt, admMdlwr)
	nurse.DELETE("/:id", nurCtrl.Delete, jwt, admMdlwr)

	// Outpatient Sessions
	outpatientSession := v1.Group("/outpatient_sessions")
	outpatientSession.GET("", outpatientSessionCtrl.GetAll, jwt)
	outpatientSession.GET("/:id", outpatientSessionCtrl.GetById, jwt)
	outpatientSession.GET("/patient/:patient_id", outpatientSessionCtrl.GetByPatientId, jwt)
	outpatientSession.GET("/doctor/:doctor_id", outpatientSessionCtrl.GetByDoctorId, jwt)
	outpatientSession.GET("/doctor/:doctor_id/unprocesseds", outpatientSessionCtrl.GetUnprocessedByDoctorId, jwt)
	outpatientSession.GET("/doctor/:doctor_id/processeds", outpatientSessionCtrl.GetProcessedByDoctorId, jwt)
	outpatientSession.POST("", outpatientSessionCtrl.Create, jwt, admMdlwr)
	outpatientSession.PUT("/:id", outpatientSessionCtrl.Update, jwt, admMdlwr)
	outpatientSession.PUT("/:id/approval", outpatientSessionCtrl.Approval, jwt, dctrMdlwr)
	outpatientSession.DELETE("/:id", outpatientSessionCtrl.Delete, jwt, admMdlwr)

	return e
}
