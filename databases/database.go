package databases

import (
	"fmt"
	"hms-backend/configs"
	"hms-backend/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	InitDB()
	InitialMigration()
}

func InitDB() {
	configs.InitConfig()

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		configs.Cfg.DbUsername,
		configs.Cfg.DbPassword,
		configs.Cfg.DbHost,
		configs.Cfg.DbPort,
		configs.Cfg.DbName,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	err := DB.AutoMigrate(
		&models.Role{},
		&models.Religion{},
		&models.Status{},
		&models.Speciality{},
		&models.Doctor{},
		&models.Nurse{},
		&models.DoctorNurse{},
		&models.DoctorSchedule{},
		&models.Patient{},
		&models.OutpatientSession{},
		&models.Treatment{},
		&models.History{},
	)
	if err != nil {
		return
	}
}
