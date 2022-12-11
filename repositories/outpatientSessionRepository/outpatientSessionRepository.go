package outpatientSessionRepository

import (
	"fmt"
	"gorm.io/gorm"
	"hms-backend/models"
	"strconv"
	"time"
)

type OutpatientSessionRepository interface {
	GetAll() ([]models.OutpatientSession, error)
	GetDesc(limit int) ([]models.OutpatientSession, error)
	GetById(id uint) (models.OutpatientSession, error)
	GetByDoctorId(doctorId uint) ([]models.OutpatientSession, error)
	GetByPatientId(patientId uint) ([]models.OutpatientSession, error)
	GetUnprocessedByDoctorId(doctorId uint) ([]models.OutpatientSession, error)
	GetProcessedByDoctorId(doctorId uint) ([]models.OutpatientSession, error)
	GetProcessedAllByDoctorId(doctorId uint) ([]models.OutpatientSession, error)
	GetApprovedByDoctorId(doctorId uint) ([]models.OutpatientSession, error)
	GetRejectedByDoctorId(doctorId uint) ([]models.OutpatientSession, error)
	GetByDate(date time.Time) ([]models.OutpatientSession, error)
	Create(user models.OutpatientSession) (models.OutpatientSession, error)
	Update(id uint, user models.OutpatientSession) (models.OutpatientSession, error)
	Approval(id uint, isApproved int) (models.OutpatientSession, error)
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

	if err := rep.db.Model([]models.OutpatientSession{}).Where("doctor_id = ? AND is_finish = ? AND is_approved = 0", doctorId, false).Find(&outpatientSessions).Error; err != nil {
		return outpatientSessions, err
	}

	return outpatientSessions, nil
}
func (rep *outpatientSessionRepository) GetProcessedByDoctorId(doctorId uint) ([]models.OutpatientSession, error) {
	var outpatientSessions []models.OutpatientSession

	if err := rep.db.Model([]models.OutpatientSession{}).Where("doctor_id = ? AND is_finish = ? AND is_approved != 0", doctorId, false).Find(&outpatientSessions).Error; err != nil {
		return outpatientSessions, err
	}

	return outpatientSessions, nil
}
func (rep *outpatientSessionRepository) GetProcessedAllByDoctorId(doctorId uint) ([]models.OutpatientSession, error) {
	var outpatientSessions []models.OutpatientSession

	if err := rep.db.Model([]models.OutpatientSession{}).Where("doctor_id = ? AND is_approved != 0", doctorId).Find(&outpatientSessions).Error; err != nil {
		return outpatientSessions, err
	}

	return outpatientSessions, nil
}
func (rep *outpatientSessionRepository) GetApprovedByDoctorId(doctorId uint) ([]models.OutpatientSession, error) {
	var outpatientSessions []models.OutpatientSession

	if err := rep.db.Model([]models.OutpatientSession{}).Where("doctor_id = ? AND is_finish = ? AND is_approved = 1", doctorId, false).Find(&outpatientSessions).Error; err != nil {
		return outpatientSessions, err
	}

	return outpatientSessions, nil
}
func (rep *outpatientSessionRepository) GetRejectedByDoctorId(doctorId uint) ([]models.OutpatientSession, error) {
	var outpatientSessions []models.OutpatientSession

	if err := rep.db.Model([]models.OutpatientSession{}).Where("doctor_id = ? AND is_approved = 2", doctorId).Find(&outpatientSessions).Error; err != nil {
		return outpatientSessions, err
	}

	return outpatientSessions, nil
}
func (rep *outpatientSessionRepository) GetByDate(date time.Time) ([]models.OutpatientSession, error) {
	var outpatientSessions []models.OutpatientSession

	dateString := strconv.Itoa(date.Year()) + "-" + strconv.Itoa(int(date.Month())) + "-"
	startTime, err := time.Parse(time.RFC3339, dateString+fmt.Sprintf("%02d", date.Day())+"T00:00:00+07:00")
	if err != nil {
		return nil, err
	}

	endTime, err := time.Parse(time.RFC3339, dateString+fmt.Sprintf("%02d", date.Day()+1)+"T00:00:00+07:00")
	if err != nil {
		return nil, err
	}

	if err := rep.db.Model([]models.OutpatientSession{}).Where("is_approved != 2 AND schedule BETWEEN ? AND ?", startTime, endTime).Find(&outpatientSessions).Error; err != nil {
		return outpatientSessions, err
	}

	return outpatientSessions, nil
}
func (rep *outpatientSessionRepository) GetDesc(limit int) ([]models.OutpatientSession, error) {
	var outpatientSessions []models.OutpatientSession

	if err := rep.db.Model([]models.OutpatientSession{}).Where("is_approved != 2").Order("schedule desc").Limit(limit).Find(&outpatientSessions).Error; err != nil {
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
func (rep *outpatientSessionRepository) Approval(id uint, isApproved int) (models.OutpatientSession, error) {
	var outpatientSession models.OutpatientSession

	err := rep.db.Model(models.OutpatientSession{}).Where("ID = ?", id).Select("is_approved").Updates(models.OutpatientSession{IsApproved: isApproved}).Error
	if err != nil {
		return outpatientSession, err
	}

	if err = rep.db.Model([]models.OutpatientSession{}).Where("ID = ?", id).First(&outpatientSession).Error; err != nil {
		return outpatientSession, err
	}

	return outpatientSession, err
}
func (rep *outpatientSessionRepository) Delete(id uint) error {
	var nurse models.OutpatientSession
	return rep.db.Delete(&nurse, id).Error
}
