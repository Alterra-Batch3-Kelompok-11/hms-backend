package doctorScheduleUseCase

import (
	"gorm.io/gorm"
	"hms-backend/constants"
	"hms-backend/dto"
	"hms-backend/models"
	"hms-backend/repositories/doctorRepository"
	"hms-backend/repositories/doctorScheduleRepository"
)

type DoctorScheduleUseCase interface {
	GetById(id uint) (dto.DoctorScheduleRes, error)
	GetByDoctorId(doctorId uint) ([]dto.DoctorScheduleRes, error)
	GetByLicenseNumber(licenseNumber string) ([]dto.DoctorScheduleRes, error)
	Create(payload dto.DoctorScheduleReq) (dto.DoctorScheduleRes, error)
	Update(id uint, payload dto.DoctorScheduleReq) (dto.DoctorScheduleRes, error)
	Delete(id uint) error
}

type doctorScheduleUseCase struct {
	doctorRep doctorRepository.DoctorRepository
	schedRep  doctorScheduleRepository.DoctorScheduleRepository
}

func New(
	dctRep doctorRepository.DoctorRepository,
	scdRep doctorScheduleRepository.DoctorScheduleRepository,
) *doctorScheduleUseCase {
	return &doctorScheduleUseCase{dctRep, scdRep}
}
func (uc *doctorScheduleUseCase) GetById(id uint) (dto.DoctorScheduleRes, error) {
	var res dto.DoctorScheduleRes
	sched, err := uc.schedRep.GetById(id)
	if err != nil {
		return res, err
	}

	res = dto.DoctorScheduleRes{
		ID:        sched.ID,
		CreatedAt: sched.CreatedAt,
		UpdatedAt: sched.UpdatedAt,
		DeletedAt: sched.DeletedAt,
		DoctorId:  sched.DoctorId,
		DayInt:    sched.Day,
		DayString: constants.Hari[sched.Day],
		StartTime: sched.StartTime,
		EndTime:   sched.EndTime,
	}

	return res, nil
}
func (uc *doctorScheduleUseCase) GetByDoctorId(doctorId uint) ([]dto.DoctorScheduleRes, error) {
	var res []dto.DoctorScheduleRes
	scheds, err := uc.schedRep.GetByDoctorId(doctorId)
	if err != nil {
		return res, err
	}

	for _, sched := range scheds {

		res = append(res, dto.DoctorScheduleRes{
			ID:        sched.ID,
			CreatedAt: sched.CreatedAt,
			UpdatedAt: sched.UpdatedAt,
			DeletedAt: sched.DeletedAt,
			DoctorId:  sched.DoctorId,
			DayInt:    sched.Day,
			DayString: constants.Hari[sched.Day],
			StartTime: sched.StartTime,
			EndTime:   sched.EndTime,
		})
	}

	return res, nil
}
func (uc *doctorScheduleUseCase) GetByLicenseNumber(licenseNumber string) ([]dto.DoctorScheduleRes, error) {
	var res []dto.DoctorScheduleRes

	doctor, err := uc.doctorRep.GetByLicenseNumber(licenseNumber)
	if err != nil {
		return res, err
	}

	scheds, err := uc.schedRep.GetByDoctorId(doctor.ID)
	if err != nil {
		return res, err
	}

	for _, sched := range scheds {

		res = append(res, dto.DoctorScheduleRes{
			ID:        sched.ID,
			CreatedAt: sched.CreatedAt,
			UpdatedAt: sched.UpdatedAt,
			DeletedAt: sched.DeletedAt,
			DoctorId:  sched.DoctorId,
			DayInt:    sched.Day,
			DayString: constants.Hari[sched.Day],
			StartTime: sched.StartTime,
			EndTime:   sched.EndTime,
		})
	}

	return res, nil
}
func (uc *doctorScheduleUseCase) Create(payload dto.DoctorScheduleReq) (dto.DoctorScheduleRes, error) {
	var res dto.DoctorScheduleRes

	sched := models.DoctorSchedule{
		Model:     gorm.Model{},
		DoctorId:  payload.DoctorId,
		Day:       payload.DayInt,
		StartTime: payload.StartTime,
		EndTime:   payload.EndTime,
	}

	resUc, err := uc.schedRep.Create(sched)
	if err != nil {
		return res, err
	}

	res = dto.DoctorScheduleRes{
		ID:        resUc.ID,
		CreatedAt: resUc.CreatedAt,
		UpdatedAt: resUc.UpdatedAt,
		DeletedAt: resUc.DeletedAt,
		DoctorId:  resUc.DoctorId,
		DayInt:    resUc.Day,
		DayString: constants.Hari[resUc.Day],
		StartTime: resUc.StartTime,
		EndTime:   resUc.EndTime,
	}

	return res, nil
}
func (uc *doctorScheduleUseCase) Update(id uint, payload dto.DoctorScheduleReq) (dto.DoctorScheduleRes, error) {
	var res dto.DoctorScheduleRes

	sched := models.DoctorSchedule{
		Model:     gorm.Model{},
		DoctorId:  payload.DoctorId,
		Day:       payload.DayInt,
		StartTime: payload.StartTime,
		EndTime:   payload.EndTime,
	}

	resUc, err := uc.schedRep.Update(id, sched)
	if err != nil {
		return res, err
	}

	res = dto.DoctorScheduleRes{
		ID:        resUc.ID,
		CreatedAt: resUc.CreatedAt,
		UpdatedAt: resUc.UpdatedAt,
		DeletedAt: resUc.DeletedAt,
		DoctorId:  resUc.DoctorId,
		DayInt:    resUc.Day,
		DayString: constants.Hari[resUc.Day],
		StartTime: resUc.StartTime,
		EndTime:   resUc.EndTime,
	}

	return res, nil
}
func (uc *doctorScheduleUseCase) Delete(id uint) error {

	err := uc.schedRep.Delete(id)
	if err != nil {
		return err
	}

	return nil
}