package doctorUseCase

import (
	"errors"
	"gorm.io/gorm"
	"hms-backend/constants"
	"hms-backend/dto"
	"hms-backend/helpers"
	"hms-backend/models"
	"hms-backend/repositories/doctorRepository"
	"hms-backend/repositories/doctorScheduleRepository"
	"hms-backend/repositories/specialityRepository"
	"hms-backend/repositories/userRepository"
	"time"
)

type DoctorUseCase interface {
	GetAll() ([]dto.DoctorRes, error)
	GetById(id uint) (dto.DoctorRes, error)
	GetByLicenseNumber(licenseNumber string) (dto.DoctorRes, error)
	GetBySpecialityId(specialityId uint) ([]dto.DoctorRes, error)
	GetToday() ([]dto.DoctorRes, error)
	Create(payload dto.UserReq) (dto.DoctorRes, error)
	Update(id uint, payload dto.UserReq) (dto.DoctorRes, error)
	Delete(id uint) error
}

type doctorUseCase struct {
	doctorRep doctorRepository.DoctorRepository
	userRep   userRepository.UserRepository
	spcRep    specialityRepository.SpecialityRepository
	scdRep    doctorScheduleRepository.DoctorScheduleRepository
}

func New(
	dctRep doctorRepository.DoctorRepository,
	usrRep userRepository.UserRepository,
	spcRep specialityRepository.SpecialityRepository,
	scdRep doctorScheduleRepository.DoctorScheduleRepository,
) *doctorUseCase {
	return &doctorUseCase{dctRep, usrRep, spcRep, scdRep}
}

func (uc *doctorUseCase) GetAll() ([]dto.DoctorRes, error) {
	var res []dto.DoctorRes
	doctors, err := uc.doctorRep.GetAll()
	if err != nil {
		return res, err
	}

	for _, doctor := range doctors {
		user, err := uc.userRep.GetById(doctor.UserId)
		if err != nil {
			return res, err
		}

		speciality, err := uc.spcRep.GetById(doctor.SpecialityId)
		if err != nil {
			return res, err
		}

		var schedules []dto.DoctorScheduleRes
		scheds, err := uc.scdRep.GetByDoctorId(doctor.ID)
		if err != nil {
			return res, err
		}

		for _, sched := range scheds {

			schedules = append(schedules, dto.DoctorScheduleRes{
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

		res = append(res, dto.DoctorRes{
			ID:              doctor.ID,
			CreatedAt:       doctor.CreatedAt,
			UpdatedAt:       doctor.UpdatedAt,
			DeletedAt:       doctor.DeletedAt,
			Name:            user.Name,
			SpecialityId:    doctor.SpecialityId,
			LicenseNumber:   doctor.LicenseNumber,
			SpecialityName:  speciality.Name,
			DoctorSchedules: schedules,
		})
	}

	return res, nil
}
func (uc *doctorUseCase) GetById(id uint) (dto.DoctorRes, error) {
	var res dto.DoctorRes
	doctor, err := uc.doctorRep.GetById(id)
	if err != nil {
		return res, err
	}
	user, err := uc.userRep.GetById(doctor.UserId)
	if err != nil {
		return res, err
	}

	speciality, err := uc.spcRep.GetById(doctor.SpecialityId)
	if err != nil {
		return res, err
	}

	var schedules []dto.DoctorScheduleRes
	scheds, err := uc.scdRep.GetByDoctorId(doctor.ID)
	if err != nil {
		return res, err
	}

	for _, sched := range scheds {

		schedules = append(schedules, dto.DoctorScheduleRes{
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

	res = dto.DoctorRes{
		ID:              doctor.ID,
		CreatedAt:       doctor.CreatedAt,
		UpdatedAt:       doctor.UpdatedAt,
		DeletedAt:       doctor.DeletedAt,
		Name:            user.Name,
		SpecialityId:    doctor.SpecialityId,
		LicenseNumber:   doctor.LicenseNumber,
		SpecialityName:  speciality.Name,
		DoctorSchedules: schedules,
	}

	return res, nil
}
func (uc *doctorUseCase) GetByLicenseNumber(licenseNumber string) (dto.DoctorRes, error) {
	var res dto.DoctorRes
	doctor, err := uc.doctorRep.GetByLicenseNumber(licenseNumber)
	if err != nil {
		return res, err
	}
	user, err := uc.userRep.GetById(doctor.UserId)
	if err != nil {
		return res, err
	}

	speciality, err := uc.spcRep.GetById(doctor.SpecialityId)
	if err != nil {
		return res, err
	}

	var schedules []dto.DoctorScheduleRes
	scheds, err := uc.scdRep.GetByDoctorId(doctor.ID)
	if err != nil {
		return res, err
	}

	for _, sched := range scheds {

		schedules = append(schedules, dto.DoctorScheduleRes{
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

	res = dto.DoctorRes{
		ID:              doctor.ID,
		CreatedAt:       doctor.CreatedAt,
		UpdatedAt:       doctor.UpdatedAt,
		DeletedAt:       doctor.DeletedAt,
		Name:            user.Name,
		SpecialityId:    doctor.SpecialityId,
		LicenseNumber:   doctor.LicenseNumber,
		SpecialityName:  speciality.Name,
		DoctorSchedules: schedules,
	}

	return res, nil
}
func (uc *doctorUseCase) GetBySpecialityId(specialityId uint) ([]dto.DoctorRes, error) {
	var res []dto.DoctorRes
	doctors, err := uc.doctorRep.GetBySpecialityId(specialityId)
	if err != nil {
		return res, err
	}

	for _, doctor := range doctors {
		user, err := uc.userRep.GetById(doctor.UserId)
		if err != nil {
			return res, err
		}

		speciality, err := uc.spcRep.GetById(doctor.SpecialityId)
		if err != nil {
			return res, err
		}

		var schedules []dto.DoctorScheduleRes
		scheds, err := uc.scdRep.GetByDoctorId(doctor.ID)
		if err != nil {
			return res, err
		}

		for _, sched := range scheds {

			schedules = append(schedules, dto.DoctorScheduleRes{
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

		res = append(res, dto.DoctorRes{
			ID:              doctor.ID,
			CreatedAt:       doctor.CreatedAt,
			UpdatedAt:       doctor.UpdatedAt,
			DeletedAt:       doctor.DeletedAt,
			Name:            user.Name,
			SpecialityId:    doctor.SpecialityId,
			LicenseNumber:   doctor.LicenseNumber,
			SpecialityName:  speciality.Name,
			DoctorSchedules: schedules,
		})
	}

	return res, nil
}
func (uc *doctorUseCase) GetToday() ([]dto.DoctorRes, error) {
	var res []dto.DoctorRes

	jakartaTime, err := helpers.TimeIn(time.Now(), "Asia/Bangkok")
	today := jakartaTime.Weekday()

	todayScheds, err := uc.scdRep.GetByDay(int(today))
	if err != nil {
		return res, err
	}

	for _, todaySched := range todayScheds {
		doctor, err := uc.doctorRep.GetById(todaySched.DoctorId)
		if err != nil {
			return res, err
		}

		user, err := uc.userRep.GetById(doctor.UserId)
		if err != nil {
			return res, err
		}

		speciality, err := uc.spcRep.GetById(doctor.SpecialityId)
		if err != nil {
			return res, err
		}

		var schedules []dto.DoctorScheduleRes
		scheds, err := uc.scdRep.GetByDoctorId(doctor.ID)
		if err != nil {
			return res, err
		}

		for _, sched := range scheds {

			schedules = append(schedules, dto.DoctorScheduleRes{
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

		res = append(res, dto.DoctorRes{
			ID:              doctor.ID,
			CreatedAt:       doctor.CreatedAt,
			UpdatedAt:       doctor.UpdatedAt,
			DeletedAt:       doctor.DeletedAt,
			Name:            user.Name,
			SpecialityId:    doctor.SpecialityId,
			LicenseNumber:   doctor.LicenseNumber,
			SpecialityName:  speciality.Name,
			DoctorSchedules: schedules,
		})
	}

	return res, nil
}
func (uc *doctorUseCase) Create(payload dto.UserReq) (dto.DoctorRes, error) {
	roleId := 2 // role id doctor

	hashedPass, err := helpers.HashPassword(payload.Password)
	if err != nil {
		return dto.DoctorRes{}, err
	}

	// TODO Check License Number if Exist
	existDoctor, _ := uc.doctorRep.GetByLicenseNumber(payload.LicenseNumber)
	if existDoctor.ID != 0 {
		return dto.DoctorRes{}, errors.New("license number already exist")
	}
	payload.Username = payload.LicenseNumber
	payload.RoleID = uint(roleId)

	// TODO check username exist
	usernameExist, _ := uc.userRep.GetByUsername(payload.Username)
	if usernameExist.ID != 0 {
		return dto.DoctorRes{}, errors.New("username already taken")
	}

	user := models.User{
		Model:    gorm.Model{},
		RoleId:   payload.RoleID,
		Username: payload.Username,
		Password: hashedPass,
		Name:     payload.Name,
	}

	resCreateUsr, err := uc.userRep.Create(user)
	if err != nil {
		return dto.DoctorRes{}, err
	}

	doctor := models.Doctor{
		Model:         gorm.Model{},
		UserId:        resCreateUsr.ID,
		SpecialityId:  payload.SpecialityID,
		LicenseNumber: payload.LicenseNumber,
	}

	resCreateDtr, err := uc.doctorRep.Create(doctor)
	if err != nil {
		return dto.DoctorRes{}, err
	}

	speciality, _ := uc.spcRep.GetById(resCreateDtr.SpecialityId)

	res := dto.DoctorRes{
		ID:              resCreateDtr.ID,
		CreatedAt:       resCreateDtr.CreatedAt,
		UpdatedAt:       resCreateDtr.UpdatedAt,
		DeletedAt:       resCreateDtr.DeletedAt,
		Name:            resCreateUsr.Name,
		SpecialityId:    resCreateDtr.SpecialityId,
		LicenseNumber:   resCreateDtr.LicenseNumber,
		SpecialityName:  speciality.Name,
		DoctorSchedules: []dto.DoctorScheduleRes{},
	}

	return res, err
}
func (uc *doctorUseCase) Update(id uint, payload dto.UserReq) (dto.DoctorRes, error) {
	roleId := 2 // role id doctor

	doctor, err := uc.doctorRep.GetById(id)
	if err != nil {
		return dto.DoctorRes{}, err
	}

	hashedPass, err := helpers.HashPassword(payload.Password)
	if err != nil {
		return dto.DoctorRes{}, err
	}

	// TODO Check License Number if Exist
	existDoctor, _ := uc.doctorRep.GetByLicenseNumberOther(payload.LicenseNumber, id)
	if existDoctor.ID != 0 {
		return dto.DoctorRes{}, errors.New("license number already exist")
	}
	payload.Username = payload.LicenseNumber
	payload.RoleID = uint(roleId)

	// TODO check username exist
	usernameExist, _ := uc.userRep.GetByUsername(payload.Username)
	if usernameExist.ID != 0 {
		return dto.DoctorRes{}, errors.New("username already taken")
	}

	user := models.User{
		Model:    gorm.Model{},
		RoleId:   payload.RoleID,
		Username: payload.Username,
		Password: hashedPass,
		Name:     payload.Name,
	}

	resUpdtUsr, err := uc.userRep.Update(id, user)
	if err != nil {
		return dto.DoctorRes{}, err
	}

	doctor = models.Doctor{
		Model:         gorm.Model{},
		UserId:        doctor.UserId,
		SpecialityId:  payload.SpecialityID,
		LicenseNumber: payload.LicenseNumber,
	}

	resUpdtDtr, err := uc.doctorRep.Update(id, doctor)
	if err != nil {
		return dto.DoctorRes{}, err
	}

	speciality, _ := uc.spcRep.GetById(resUpdtDtr.SpecialityId)

	var schedules []dto.DoctorScheduleRes
	scheds, err := uc.scdRep.GetByDoctorId(doctor.ID)
	if err != nil {
		return dto.DoctorRes{}, err
	}

	for _, sched := range scheds {

		schedules = append(schedules, dto.DoctorScheduleRes{
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

	res := dto.DoctorRes{
		ID:              resUpdtDtr.ID,
		CreatedAt:       resUpdtDtr.CreatedAt,
		UpdatedAt:       resUpdtDtr.UpdatedAt,
		DeletedAt:       resUpdtDtr.DeletedAt,
		Name:            resUpdtUsr.Name,
		SpecialityId:    resUpdtDtr.SpecialityId,
		LicenseNumber:   resUpdtDtr.LicenseNumber,
		SpecialityName:  speciality.Name,
		DoctorSchedules: schedules,
	}

	return res, err
}
func (uc *doctorUseCase) Delete(id uint) error {
	doctor, err := uc.doctorRep.GetById(id)
	if err != nil {
		return err
	}
	usrId := doctor.UserId

	err = uc.doctorRep.Delete(id)
	if err != nil {
		return err
	}

	err = uc.userRep.Delete(usrId)
	if err != nil {
		return err
	}

	return nil
}
