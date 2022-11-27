package doctorRepository

import (
	"gorm.io/gorm"
	"hms-backend/models"
)

type DoctorRepository interface {
	GetAll() ([]models.Doctor, error)
	GetById(id uint) (models.Doctor, error)
	GetByLicenseNumber(licenseNumber string) (models.Doctor, error)
	GetByLicenseNumberOther(licenseNumber string, id uint) (models.Doctor, error)
	GetBySpecialityId(specialityId uint) ([]models.Doctor, error)
	Create(user models.Doctor) (models.Doctor, error)
	Update(id uint, user models.Doctor) (models.Doctor, error)
	Delete(id uint) error
}

type doctorRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *doctorRepository {
	return &doctorRepository{db}
}

func (rep *doctorRepository) GetAll() ([]models.Doctor, error) {
	var doctors []models.Doctor

	if err := rep.db.Model([]models.Doctor{}).Find(&doctors).Error; err != nil {
		return doctors, err
	}

	return doctors, nil
}
func (rep *doctorRepository) GetById(id uint) (models.Doctor, error) {
	doctor := models.Doctor{}

	if err := rep.db.Model(models.Doctor{}).Where("ID = ?", id).First(&doctor).Error; err != nil {
		return doctor, err
	}

	return doctor, nil
}
func (rep *doctorRepository) GetByLicenseNumber(licenseNumber string) (models.Doctor, error) {
	doctor := models.Doctor{}

	if err := rep.db.Model(models.Doctor{}).Where("license_number = ?", licenseNumber).First(&doctor).Error; err != nil {
		return doctor, err
	}

	return doctor, nil
}
func (rep *doctorRepository) GetByLicenseNumberOther(licenseNumber string, id uint) (models.Doctor, error) {
	doctor := models.Doctor{}

	if err := rep.db.Model(models.Doctor{}).Where("license_number = ? AND ID != ?", licenseNumber, id).First(&doctor).Error; err != nil {
		return doctor, err
	}

	return doctor, nil
}
func (rep *doctorRepository) GetBySpecialityId(specialityId uint) ([]models.Doctor, error) {
	var doctors []models.Doctor

	if err := rep.db.Model([]models.Doctor{}).Where("speciality_id = ?", specialityId).Find(&doctors).Error; err != nil {
		return doctors, err
	}

	return doctors, nil
}
func (rep *doctorRepository) Create(doctor models.Doctor) (models.Doctor, error) {
	err := rep.db.Create(&doctor).Error
	return doctor, err
}
func (rep *doctorRepository) Update(id uint, doctor models.Doctor) (models.Doctor, error) {
	err := rep.db.Model(models.Doctor{}).Where("ID = ?", id).Updates(&doctor).Error
	return doctor, err
}
func (rep *doctorRepository) Delete(id uint) error {
	var doctor models.Doctor
	return rep.db.Delete(&doctor, id).Error
}
