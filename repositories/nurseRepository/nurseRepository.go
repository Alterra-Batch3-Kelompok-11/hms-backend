package nurseRepository

import (
	"gorm.io/gorm"
	"hms-backend/models"
)

type NurseRepository interface {
	GetAll() ([]models.Nurse, error)
	GetById(id uint) (models.Nurse, error)
	GetByUserId(userId uint) (models.Nurse, error)
	GetByLicenseNumber(licenseNumber string) (models.Nurse, error)
	Create(user models.Nurse) (models.Nurse, error)
	Update(id uint, user models.Nurse) (models.Nurse, error)
	Delete(id uint) error
}

type nurseRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *nurseRepository {
	return &nurseRepository{db}
}

func (rep *nurseRepository) GetAll() ([]models.Nurse, error) {
	var nurses []models.Nurse

	if err := rep.db.Model([]models.Nurse{}).Find(&nurses).Error; err != nil {
		return nurses, err
	}

	return nurses, nil
}
func (rep *nurseRepository) GetById(id uint) (models.Nurse, error) {
	nurse := models.Nurse{}

	if err := rep.db.Model(models.Nurse{}).Where("ID = ?", id).First(&nurse).Error; err != nil {
		return nurse, err
	}

	return nurse, nil
}
func (rep *nurseRepository) GetByUserId(userId uint) (models.Nurse, error) {
	nurse := models.Nurse{}

	if err := rep.db.Model(models.Nurse{}).Where("user_id = ?", userId).First(&nurse).Error; err != nil {
		return nurse, err
	}

	return nurse, nil
}
func (rep *nurseRepository) GetByLicenseNumber(licenseNumber string) (models.Nurse, error) {
	nurse := models.Nurse{}

	if err := rep.db.Model(models.Nurse{}).Where("license_number = ?", licenseNumber).First(&nurse).Error; err != nil {
		return nurse, err
	}

	return nurse, nil
}
func (rep *nurseRepository) Create(nurse models.Nurse) (models.Nurse, error) {
	err := rep.db.Create(&nurse).Error
	return nurse, err
}
func (rep *nurseRepository) Update(id uint, nurse models.Nurse) (models.Nurse, error) {
	err := rep.db.Model(models.Nurse{}).Where("ID = ?", id).Updates(&nurse).Error
	return nurse, err
}
func (rep *nurseRepository) Delete(id uint) error {
	var nurse models.Nurse
	return rep.db.Delete(&nurse, id).Error
}
