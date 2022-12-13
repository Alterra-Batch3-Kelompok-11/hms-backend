package treatmentRepository

import (
	"gorm.io/gorm"
	"hms-backend/models"
)

type TreatmentRepository interface {
	GetAll() ([]models.Treatment, error)
	GetById(id uint) (models.Treatment, error)
	GetByOutpatientSessionId(outpatientSessionId uint) (models.Treatment, error)
	Create(payload models.Treatment) (models.Treatment, error)
	Update(id uint, payload models.Treatment) (models.Treatment, error)
	Delete(id uint) error
}

type treatmentRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *treatmentRepository {
	return &treatmentRepository{db}
}

func (rep *treatmentRepository) GetAll() ([]models.Treatment, error) {
	var treatments []models.Treatment

	if err := rep.db.Model([]models.Treatment{}).Find(&treatments).Error; err != nil {
		return treatments, err
	}

	return treatments, nil
}
func (rep *treatmentRepository) GetById(id uint) (models.Treatment, error) {
	treatment := models.Treatment{}

	if err := rep.db.Model(models.Treatment{}).Where("ID = ?", id).First(&treatment).Error; err != nil {
		return treatment, err
	}

	return treatment, nil
}
func (rep *treatmentRepository) GetByOutpatientSessionId(outpatientSessionId uint) (models.Treatment, error) {
	treatment := models.Treatment{}

	if err := rep.db.Model(models.Treatment{}).Where("session_id = ?", outpatientSessionId).First(&treatment).Error; err != nil {
		return treatment, err
	}

	return treatment, nil
}
func (rep *treatmentRepository) Create(payload models.Treatment) (models.Treatment, error) {
	err := rep.db.Create(&payload).Error
	return payload, err
}
func (rep *treatmentRepository) Update(id uint, payload models.Treatment) (models.Treatment, error) {
	err := rep.db.Model(models.Treatment{}).Where("ID = ?", id).Updates(&payload).Error
	return payload, err
}
func (rep *treatmentRepository) Delete(id uint) error {
	var treatment models.Treatment
	return rep.db.Delete(&treatment, id).Error
}
