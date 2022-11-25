package patientRepository

import (
	"hms-backend/models"

	"gorm.io/gorm"
)

type PatientRepository interface {
	GetAll() ([]models.Patient, error)
	GetById(id uint) (models.Patient, error)
	Create(payload models.Patient) (models.Patient, error)
	Update(id uint, payload models.Patient) (models.Patient, error)
	Delete(id uint) error
}

type patientRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *patientRepository {
	return &patientRepository{db}
}

func (rep *patientRepository) GetAll() ([]models.Patient, error) {
	var patient []models.Patient

	if err := rep.db.Model([]models.Patient{}).Find(&patient).Error; err != nil {
		return patient, err
	}

	return patient, nil
}

func (rep *patientRepository) GetById(id uint) (models.Patient, error) {
	patient := models.Patient{}

	if err := rep.db.Model(models.Patient{}).Where("ID = ?", id).First(&patient).Error; err != nil {
		return patient, err
	}

	return patient, nil
}

func (rep *patientRepository) Create(payload models.Patient) (models.Patient, error) {
	err := rep.db.Create(&payload).Error
	return payload, err
}

func (rep *patientRepository) Update(id uint, payload models.Patient) (models.Patient, error) {
	err := rep.db.Model(models.Patient{}).Where("ID = ?", id).Updates(&payload).Error
	return payload, err
}

func (rep *patientRepository) Delete(id uint) error {
	var patient models.Patient
	return rep.db.Delete(&patient, id).Error
}
