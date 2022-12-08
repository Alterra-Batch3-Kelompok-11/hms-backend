package doctorScheduleRepository

import (
	"gorm.io/gorm"
	"hms-backend/models"
)

type DoctorScheduleRepository interface {
	GetById(id uint) (models.DoctorSchedule, error)
	GetByDoctorId(doctorId uint) ([]models.DoctorSchedule, error)
	GetByDoctorIdDay(doctorId uint, dayInt int) (models.DoctorSchedule, error)
	GetByDay(dayInt int) ([]models.DoctorSchedule, error)
	Create(sched models.DoctorSchedule) (models.DoctorSchedule, error)
	Update(id uint, sched models.DoctorSchedule) (models.DoctorSchedule, error)
	Delete(id uint) error
}

type doctorScheduleRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *doctorScheduleRepository {
	return &doctorScheduleRepository{db}
}

func (rep *doctorScheduleRepository) GetById(id uint) (models.DoctorSchedule, error) {
	sched := models.DoctorSchedule{}

	if err := rep.db.Model(models.DoctorSchedule{}).Where("ID = ?", id).First(&sched).Error; err != nil {
		return sched, err
	}

	return sched, nil
}
func (rep *doctorScheduleRepository) GetByDoctorId(doctorId uint) ([]models.DoctorSchedule, error) {
	var scheds []models.DoctorSchedule

	if err := rep.db.Model(models.DoctorSchedule{}).Where("doctor_id = ?", doctorId).Order("day").Order("start_time").Find(&scheds).Error; err != nil {
		return scheds, err
	}

	return scheds, nil
}
func (rep *doctorScheduleRepository) GetByDoctorIdDay(doctorId uint, dayInt int) (models.DoctorSchedule, error) {
	var scheds models.DoctorSchedule

	if err := rep.db.Model(models.DoctorSchedule{}).Where("doctor_id = ? AND day = ?", doctorId, dayInt).Last(&scheds).Error; err != nil {
		return scheds, err
	}

	return scheds, nil
}
func (rep *doctorScheduleRepository) GetByDay(dayInt int) ([]models.DoctorSchedule, error) {
	var scheds []models.DoctorSchedule

	if err := rep.db.Model(models.DoctorSchedule{}).Where("day = ?", dayInt).Order("day").Order("start_time").Find(&scheds).Error; err != nil {
		return scheds, err
	}

	return scheds, nil
}
func (rep *doctorScheduleRepository) Create(doctor models.DoctorSchedule) (models.DoctorSchedule, error) {
	err := rep.db.Create(&doctor).Error
	return doctor, err
}
func (rep *doctorScheduleRepository) Update(id uint, sched models.DoctorSchedule) (models.DoctorSchedule, error) {
	err := rep.db.Model(models.DoctorSchedule{}).Where("ID = ?", id).Updates(&sched).Error
	return sched, err
}
func (rep *doctorScheduleRepository) Delete(id uint) error {
	var sched models.DoctorSchedule
	return rep.db.Delete(&sched, id).Error
}
