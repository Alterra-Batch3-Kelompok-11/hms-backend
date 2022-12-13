package historyRepository

import (
	"gorm.io/gorm"
	"hms-backend/models"
)

type HistoryRepository interface {
	GetAll() ([]models.History, error)
	GetById(id uint) (models.History, error)
	GetByPatientId(patientId uint) ([]models.History, error)
	GetByDoctorId(doctorId uint) ([]models.History, error)
	Create(payload models.History) (models.History, error)
}

type historyRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *historyRepository {
	return &historyRepository{db}
}

func (rep *historyRepository) GetAll() ([]models.History, error) {
	var histories []models.History

	if err := rep.db.Model([]models.History{}).Find(&histories).Error; err != nil {
		return histories, err
	}

	return histories, nil
}
func (rep *historyRepository) GetById(id uint) (models.History, error) {
	history := models.History{}

	if err := rep.db.Model(models.History{}).Where("ID = ?", id).First(&history).Error; err != nil {
		return history, err
	}

	return history, nil
}
func (rep *historyRepository) GetByPatientId(patientId uint) ([]models.History, error) {
	var histories []models.History

	if err := rep.db.Model(models.Treatment{}).Where("patient_id = ?", patientId).Find(&histories).Error; err != nil {
		return histories, err
	}

	return histories, nil
}
func (rep *historyRepository) GetByDoctorId(doctorId uint) ([]models.History, error) {
	var histories []models.History

	if err := rep.db.Model(models.Treatment{}).Where("doctor_id = ?", doctorId).Find(&histories).Error; err != nil {
		return histories, err
	}

	return histories, nil
}
func (rep *historyRepository) Create(payload models.History) (models.History, error) {
	err := rep.db.Create(&payload).Error
	return payload, err
}
