package outpatientSessionRepository

import (
	"gorm.io/gorm"
	"hms-backend/models"
)

type OutpatientSessionRepository interface {
	GetAll() ([]models.OutpatientSession, error)
	GetById(id uint) (models.OutpatientSession, error)
	GetByDoctorId(doctorId uint) ([]models.OutpatientSession, error)
	GetByPatientId(patientId uint) ([]models.OutpatientSession, error)
	GetUnprocessedByDoctorId(doctorId uint) ([]models.OutpatientSession, error)
	GetProcessedByDoctorId(doctorId uint) ([]models.OutpatientSession, error)
	Create(user models.OutpatientSession) (models.OutpatientSession, error)
	Update(id uint, user models.OutpatientSession) (models.OutpatientSession, error)
	Delete(id uint) error
}

type outpatientSessionRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *outpatientSessionRepository {
	return &outpatientSessionRepository{db}
}

func (rep *outpatientSessionRepository) GetAll() ([]models.OutpatientSession, error) {
	var outpatientSessions []models.OutpatientSession

	if err := rep.db.Model([]models.OutpatientSession{}).Find(&outpatientSessions).Error; err != nil {
		return outpatientSessions, err
	}

	return outpatientSessions, nil
}
func (rep *outpatientSessionRepository) GetById(id uint) (models.OutpatientSession, error) {
	var outpatientSession models.OutpatientSession

	if err := rep.db.Model([]models.OutpatientSession{}).Where("ID = ?", id).First(&outpatientSession).Error; err != nil {
		return outpatientSession, err
	}

	return outpatientSession, nil
}
func (rep *outpatientSessionRepository) GetByDoctorId(doctorId uint) ([]models.OutpatientSession, error) {
	var outpatientSessions []models.OutpatientSession

	if err := rep.db.Model([]models.OutpatientSession{}).Where("doctor_id = ?", doctorId).Find(&outpatientSessions).Error; err != nil {
		return outpatientSessions, err
	}

	return outpatientSessions, nil
}
func (rep *outpatientSessionRepository) GetByPatientId(patientId uint) ([]models.OutpatientSession, error) {
	var outpatientSessions []models.OutpatientSession

	if err := rep.db.Model([]models.OutpatientSession{}).Where("patient_id = ?", patientId).Find(&outpatientSessions).Error; err != nil {
		return outpatientSessions, err
	}

	return outpatientSessions, nil
}
func (rep *outpatientSessionRepository) GetUnprocessedByDoctorId(doctorId uint) ([]models.OutpatientSession, error) {
	var outpatientSessions []models.OutpatientSession

	if err := rep.db.Model([]models.OutpatientSession{}).Where("doctor_id = ? AND is_approved = 0", doctorId).Find(&outpatientSessions).Error; err != nil {
		return outpatientSessions, err
	}

	return outpatientSessions, nil
}
func (rep *outpatientSessionRepository) GetProcessedByDoctorId(doctorId uint) ([]models.OutpatientSession, error) {
	var outpatientSessions []models.OutpatientSession

	if err := rep.db.Model([]models.OutpatientSession{}).Where("doctor_id = ? AND is_approved != 0", doctorId).Find(&outpatientSessions).Error; err != nil {
		return outpatientSessions, err
	}

	return outpatientSessions, nil
}
func (rep *outpatientSessionRepository) Create(outpatientSession models.OutpatientSession) (models.OutpatientSession, error) {
	err := rep.db.Create(&outpatientSession).Error
	return outpatientSession, err
}
func (rep *outpatientSessionRepository) Update(id uint, outpatientSession models.OutpatientSession) (models.OutpatientSession, error) {
	err := rep.db.Model(models.OutpatientSession{}).Where("ID = ?", id).Updates(&outpatientSession).Error
	return outpatientSession, err
}
func (rep *outpatientSessionRepository) Delete(id uint) error {
	var nurse models.OutpatientSession
	return rep.db.Delete(&nurse, id).Error
}
