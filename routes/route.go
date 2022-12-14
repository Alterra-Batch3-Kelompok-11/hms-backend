package routes

import (
	"hms-backend/configs"
	"hms-backend/controllers/authController"
	"hms-backend/controllers/dashboardController"
	"hms-backend/controllers/doctorController"
	"hms-backend/controllers/doctorScheduleController"
	"hms-backend/controllers/historyController"
	"hms-backend/controllers/nurseController"
	"hms-backend/controllers/outpatientSessionController"
	"hms-backend/controllers/patientConditionController"
	"hms-backend/controllers/patientController"
	"hms-backend/controllers/religionController"
	"hms-backend/controllers/roleController"
	"hms-backend/controllers/specialityController"
	"hms-backend/middlewares"
	"hms-backend/repositories/doctorRepository"
	"hms-backend/repositories/doctorScheduleRepository"
	"hms-backend/repositories/historyRepository"
	"hms-backend/repositories/nurseRepository"
	"hms-backend/repositories/outpatientSessionRepository"
	"hms-backend/repositories/patientRepository"
	"hms-backend/repositories/religionRepository"
	"hms-backend/repositories/roleRepository"
	"hms-backend/repositories/specialityRepository"
	"hms-backend/repositories/treatmentRepository"
	"hms-backend/repositories/userRepository"
	"hms-backend/usecases/authUseCase"
	"hms-backend/usecases/dashboardUseCase"
	"hms-backend/usecases/doctorScheduleUseCase"
	"hms-backend/usecases/doctorUseCase"
	"hms-backend/usecases/historyUseCase"
	"hms-backend/usecases/nurseUseCase"
	"hms-backend/usecases/outpatientSessionUseCase"
	"hms-backend/usecases/patientConditionUseCase"
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
	treatmentRepo := treatmentRepository.New(db)
	historyRepo := historyRepository.New(db)

	// Use Cases
	authUc := authUseCase.New(usrRepo, dtrRepo, nrsRepo)
	rlUc := roleUseCase.New(rlRepo)
	spcUc := specialityUseCase.New(spcRepo)
	patUc := patientUseCase.New(patRepo)
	rlgUc := religionUseCase.New(rlgRepo)
	dtrUc := doctorUseCase.New(dtrRepo, usrRepo, spcRepo, dtrSchedRepo)
	dtrSchdUc := doctorScheduleUseCase.New(dtrRepo, dtrSchedRepo)
	nurUC := nurseUseCase.New(nurRepo, usrRepo, spcRepo)
	outPatientSessionUC := outpatientSessionUseCase.New(outPatientSessionRepo, usrRepo, dtrRepo, spcRepo, dtrSchedRepo, patRepo)
	dashboardUC := dashboardUseCase.New(outPatientSessionRepo, usrRepo, dtrRepo, spcRepo, dtrSchedRepo, nurRepo, patRepo, rlgRepo)
	patientConditionUC := patientConditionUseCase.New(treatmentRepo, outPatientSessionRepo, usrRepo, dtrRepo, spcRepo, dtrSchedRepo, patRepo, historyRepo)
	historyUC := historyUseCase.New(outPatientSessionRepo, patRepo)

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
	dashboardCtrl := dashboardController.New(dashboardUC)
	patientConditionCtrl := patientConditionController.New(patientConditionUC)
	historyCtrl := historyController.New(historyUC)

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
	v1.GET("/auth/refresh", authCtrl.RefreshToken, jwt)

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
	outpatientSession.GET("", outpatientSessionCtrl.GetAll)
	outpatientSession.GET("/:id", outpatientSessionCtrl.GetById)
	outpatientSession.GET("/patient/:patient_id", outpatientSessionCtrl.GetByPatientId)
	outpatientSession.GET("/doctor/:doctor_id", outpatientSessionCtrl.GetByDoctorId)
	outpatientSession.GET("/doctor/:doctor_id/unprocesseds", outpatientSessionCtrl.GetUnprocessedByDoctorId)
	outpatientSession.GET("/doctor/:doctor_id/processeds", outpatientSessionCtrl.GetProcessedByDoctorId)
	outpatientSession.GET("/doctor/:doctor_id/approveds", outpatientSessionCtrl.GetApprovedByDoctorId)
	outpatientSession.GET("/doctor/:doctor_id/rejecteds", outpatientSessionCtrl.GetRejectedByDoctorId)
	outpatientSession.POST("", outpatientSessionCtrl.Create, jwt, admMdlwr)
	outpatientSession.PUT("/:id", outpatientSessionCtrl.Update, jwt, admMdlwr)
	outpatientSession.PUT("/:id/approval", outpatientSessionCtrl.Approval, jwt, dctrMdlwr)
	outpatientSession.DELETE("/:id", outpatientSessionCtrl.Delete, jwt, admMdlwr)

	// Patient Conditions / Treatments
	patientCondition := v1.Group("/patient_conditions")
	patientCondition.GET("", patientConditionCtrl.GetAll)
	patientCondition.GET("/:id", patientConditionCtrl.GetById)
	patientCondition.GET("/patient/:patient_id", patientConditionCtrl.GetByPatientId)
	patientCondition.GET("/doctor/:doctor_id", patientConditionCtrl.GetByDoctorId)
	patientCondition.POST("", patientConditionCtrl.Create, jwt)

	// Histories
	history := v1.Group("/histories")
	history.GET("/doctor/:doctor_id/outpatient_sessions", historyCtrl.GetOutpatientSessionHistory)
	history.GET("/doctor/:doctor_id/approvals", historyCtrl.GetApprovalHistory)

	// For Dashboard
	dashboard := v1.Group("/dashboard")
	dashboard.GET("/web", dashboardCtrl.GetDataDashboardWeb)
	dashboard.GET("/mobile/doctor/:doctor_id", dashboardCtrl.GetDataDashboardMobile)

	return e
}
