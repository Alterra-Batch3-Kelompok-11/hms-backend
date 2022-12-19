package nurseUseCase

import (
	"errors"
	"fmt"
	"hms-backend/constants"
	"hms-backend/dto"
	"hms-backend/helpers"
	"hms-backend/models"
	"hms-backend/repositories/nurseRepository"
	"hms-backend/repositories/specialityRepository"
	"hms-backend/repositories/userRepository"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

type NurseUseCase interface {
	GetAll() ([]dto.NurseRes, error)
	GetById(id uint) (dto.NurseRes, error)
	GetByLicenseNumber(licenseNumber string) (dto.NurseRes, error)
	Create(payload dto.NurseReq) (dto.NurseRes, error)
	Update(id uint, payload dto.NurseReq) (dto.NurseRes, error)
	Delete(id uint) error
}

type nurseUseCase struct {
	nurseRep nurseRepository.NurseRepository
	userRep  userRepository.UserRepository
	spcRep   specialityRepository.SpecialityRepository
}

func New(
	nurseRep nurseRepository.NurseRepository,
	usrRep userRepository.UserRepository,
	spcRep specialityRepository.SpecialityRepository,
) *nurseUseCase {
	return &nurseUseCase{nurseRep, usrRep, spcRep}
}

func (uc *nurseUseCase) GetAll() ([]dto.NurseRes, error) {
	var res []dto.NurseRes
	nurses, err := uc.nurseRep.GetAll()
	if err != nil {
		return res, err
	}

	for _, nurse := range nurses {
		birthDateString := strconv.Itoa(nurse.BirthDate.Year()) + "-" + strconv.Itoa(int(nurse.BirthDate.Month())) + "-" + fmt.Sprintf("%02d", nurse.BirthDate.Day())

		birthDateIndoString := fmt.Sprintf("%02d", nurse.BirthDate.Day()) + " " +
			constants.Bulan[int(nurse.BirthDate.Month())] + " " +
			strconv.Itoa(nurse.BirthDate.Year())

		speciality, err := uc.spcRep.GetById(nurse.SpecialityId)
		if err != nil {
			return res, err
		}

		user, err := uc.userRep.GetById(nurse.UserId)
		if err != nil {
			return res, err
		}

		res = append(res, dto.NurseRes{
			ID:                  nurse.ID,
			UserID:              nurse.UserId,
			DoctorID:            nurse.DoctorId,
			Name:                user.Name,
			SpecialityId:        nurse.SpecialityId,
			LicenseNumber:       nurse.LicenseNumber,
			SpecialityName:      speciality.Name,
			BirthDate:           nurse.BirthDate,
			BirthDateString:     birthDateString,
			BirthDateStringIndo: birthDateIndoString,
			ProfilePic:          nurse.ProfilePic,
			Phone:               nurse.Phone,
			MaritalStatus:       nurse.MaritalStatus,
			Email:               nurse.Email,
			Nira:                nurse.Nira,
			SIP:                 nurse.SIP,
			CreatedAt:           nurse.CreatedAt,
			UpdatedAt:           nurse.UpdatedAt,
			DeletedAt:           nurse.DeletedAt,
		})
	}

	return res, nil
}

func (uc *nurseUseCase) GetById(id uint) (dto.NurseRes, error) {
	var res dto.NurseRes
	nurse, err := uc.nurseRep.GetById(id)
	if err != nil {
		return res, err
	}

	birthDateString := strconv.Itoa(nurse.BirthDate.Year()) + "-" + strconv.Itoa(int(nurse.BirthDate.Month())) + "-" + fmt.Sprintf("%02d", nurse.BirthDate.Day())

	birthDateIndoString := fmt.Sprintf("%02d", nurse.BirthDate.Day()) + " " +
		constants.Bulan[int(nurse.BirthDate.Month())] + " " +
		strconv.Itoa(nurse.BirthDate.Year())

	speciality, err := uc.spcRep.GetById(nurse.SpecialityId)
	if err != nil {
		return res, err
	}

	user, err := uc.userRep.GetById(nurse.UserId)
	if err != nil {
		return res, err
	}

	res = dto.NurseRes{
		ID:                  nurse.ID,
		UserID:              nurse.UserId,
		DoctorID:            nurse.DoctorId,
		Name:                user.Name,
		SpecialityId:        nurse.SpecialityId,
		LicenseNumber:       nurse.LicenseNumber,
		SpecialityName:      speciality.Name,
		BirthDate:           nurse.BirthDate,
		BirthDateString:     birthDateString,
		BirthDateStringIndo: birthDateIndoString,
		ProfilePic:          nurse.ProfilePic,
		Phone:               nurse.Phone,
		MaritalStatus:       nurse.MaritalStatus,
		Email:               nurse.Email,
		Nira:                nurse.Nira,
		SIP:                 nurse.SIP,
		CreatedAt:           nurse.CreatedAt,
		UpdatedAt:           nurse.UpdatedAt,
		DeletedAt:           nurse.DeletedAt,
	}

	return res, nil
}

func (uc *nurseUseCase) GetByLicenseNumber(licenseNumber string) (dto.NurseRes, error) {
	var res dto.NurseRes
	nurse, err := uc.nurseRep.GetByLicenseNumber(licenseNumber)
	if err != nil {
		return res, err
	}

	birthDateString := strconv.Itoa(nurse.BirthDate.Year()) + "-" + strconv.Itoa(int(nurse.BirthDate.Month())) + "-" + fmt.Sprintf("%02d", nurse.BirthDate.Day())

	birthDateIndoString := fmt.Sprintf("%02d", nurse.BirthDate.Day()) + " " +
		constants.Bulan[int(nurse.BirthDate.Month())] + " " +
		strconv.Itoa(nurse.BirthDate.Year())

	speciality, err := uc.spcRep.GetById(nurse.SpecialityId)
	if err != nil {
		return res, err
	}

	user, err := uc.userRep.GetById(nurse.UserId)
	if err != nil {
		return res, err
	}

	res = dto.NurseRes{
		ID:                  nurse.ID,
		UserID:              nurse.UserId,
		DoctorID:            nurse.DoctorId,
		Name:                user.Name,
		SpecialityId:        nurse.SpecialityId,
		LicenseNumber:       nurse.LicenseNumber,
		SpecialityName:      speciality.Name,
		BirthDate:           nurse.BirthDate,
		BirthDateString:     birthDateString,
		BirthDateStringIndo: birthDateIndoString,
		ProfilePic:          nurse.ProfilePic,
		Phone:               nurse.Phone,
		MaritalStatus:       nurse.MaritalStatus,
		Email:               nurse.Email,
		Nira:                nurse.Nira,
		SIP:                 nurse.SIP,
		CreatedAt:           nurse.CreatedAt,
		UpdatedAt:           nurse.UpdatedAt,
		DeletedAt:           nurse.DeletedAt,
	}

	return res, nil
}

func (uc *nurseUseCase) Create(payload dto.NurseReq) (dto.NurseRes, error) {
	splitedBirthDate := strings.Split(payload.BirthDate[0:10], "-")

	birthDateTimeString := splitedBirthDate[2] + "-" + splitedBirthDate[1] + "-" + splitedBirthDate[0] + "T00:00:00+07:00"
	birthDateTime, err := time.Parse(time.RFC3339, birthDateTimeString)
	if err != nil {
		return dto.NurseRes{}, err
	}

	roleId := 3 // role id nurse

	hashedPass, err := helpers.HashPassword(payload.Password)
	if err != nil {
		return dto.NurseRes{}, err
	}

	// TODO Check License Number if Exist
	existNurse, _ := uc.nurseRep.GetByLicenseNumber(payload.LicenseNumber)
	if existNurse.ID != 0 {
		return dto.NurseRes{}, errors.New("license number already exist")
	}

	// TODO check username exist
	usernameExist, _ := uc.userRep.GetByUsername(payload.LicenseNumber)
	if usernameExist.ID != 0 {
		return dto.NurseRes{}, errors.New("username already taken")
	}

	user := models.User{
		Model:    gorm.Model{},
		RoleId:   uint(roleId),
		Username: payload.LicenseNumber,
		Password: hashedPass,
		Name:     payload.Name,
	}

	resCreateUsr, err := uc.userRep.Create(user)
	if err != nil {
		return dto.NurseRes{}, err
	}

	nurse := models.Nurse{
		Model:         gorm.Model{},
		UserId:        resCreateUsr.ID,
		DoctorId:      payload.DoctorID,
		LicenseNumber: payload.LicenseNumber,
		SpecialityId:  payload.SpecialityID,
		ProfilePic:    payload.ProfilePic,
		BirthDate:     birthDateTime,
		Phone:         payload.Phone,
		MaritalStatus: payload.MaritalStatus,
		Email:         payload.Email,
		Nira:          payload.Nira,
		SIP:           payload.SIP,
	}

	var res dto.NurseRes
	nurse, err = uc.nurseRep.Create(nurse)
	if err != nil {
		return res, err
	}

	birthDateString := strconv.Itoa(nurse.BirthDate.Year()) + "-" + strconv.Itoa(int(nurse.BirthDate.Month())) + "-" + fmt.Sprintf("%02d", nurse.BirthDate.Day())

	birthDateIndoString := fmt.Sprintf("%02d", nurse.BirthDate.Day()) + " " +
		constants.Bulan[int(nurse.BirthDate.Month())] + " " +
		strconv.Itoa(nurse.BirthDate.Year())

	speciality, err := uc.spcRep.GetById(nurse.SpecialityId)
	if err != nil {
		return res, err
	}

	res = dto.NurseRes{
		ID:                  nurse.ID,
		UserID:              nurse.UserId,
		DoctorID:            nurse.DoctorId,
		SpecialityId:        nurse.SpecialityId,
		LicenseNumber:       nurse.LicenseNumber,
		SpecialityName:      speciality.Name,
		BirthDate:           nurse.BirthDate,
		BirthDateString:     birthDateString,
		BirthDateStringIndo: birthDateIndoString,
		ProfilePic:          nurse.ProfilePic,
		Phone:               nurse.Phone,
		MaritalStatus:       nurse.MaritalStatus,
		Email:               nurse.Email,
		Nira:                nurse.Nira,
		SIP:                 nurse.SIP,
		CreatedAt:           nurse.CreatedAt,
		UpdatedAt:           nurse.UpdatedAt,
		DeletedAt:           nurse.DeletedAt,
	}

	return res, nil
}

func (uc *nurseUseCase) Update(id uint, payload dto.NurseReq) (dto.NurseRes, error) {
	roleId := 3 // role id nurse

	_, err := uc.nurseRep.GetById(id)
	if err != nil {
		return dto.NurseRes{}, err
	}

	hashedPass, err := helpers.HashPassword(payload.Password)
	if err != nil {
		return dto.NurseRes{}, err
	}

	// TODO Check License Number if Exist
	existNurse, _ := uc.nurseRep.GetByLicenseNumberOther(payload.LicenseNumber, id)
	if existNurse.ID != 0 {
		return dto.NurseRes{}, errors.New("license number already exist")
	}

	// TODO check username exist
	usernameExist, _ := uc.userRep.GetByUsername(payload.LicenseNumber)
	if usernameExist.ID != 0 {
		return dto.NurseRes{}, errors.New("username already taken")
	}

	user := models.User{
		Model:    gorm.Model{},
		RoleId:   uint(roleId),
		Username: payload.LicenseNumber,
		Password: hashedPass,
		Name:     payload.Name,
	}

	resUpdtUsr, err := uc.userRep.Update(id, user)
	if err != nil {
		return dto.NurseRes{}, err
	}

	splitedBirthDate := strings.Split(payload.BirthDate[0:10], "-")

	birthDateTimeString := splitedBirthDate[2] + "-" + splitedBirthDate[1] + "-" + splitedBirthDate[0] + "T00:00:00+07:00"
	birthDateTime, err := time.Parse(time.RFC3339, birthDateTimeString)
	if err != nil {
		return dto.NurseRes{}, err
	}

	nurse := models.Nurse{
		Model:         gorm.Model{},
		UserId:        resUpdtUsr.ID,
		DoctorId:      payload.DoctorID,
		LicenseNumber: payload.LicenseNumber,
		SpecialityId:  payload.SpecialityID,
		ProfilePic:    payload.ProfilePic,
		BirthDate:     birthDateTime,
		Phone:         payload.Phone,
		MaritalStatus: payload.MaritalStatus,
		Email:         payload.Email,
		Nira:          payload.Nira,
		SIP:           payload.SIP,
	}

	var res dto.NurseRes
	nurse, err = uc.nurseRep.Update(id, nurse)
	if err != nil {
		return res, err
	}

	birthDateString := fmt.Sprintf("%02d", nurse.BirthDate.Day()) + "-" + strconv.Itoa(int(nurse.BirthDate.Month())) + "-" + strconv.Itoa(nurse.BirthDate.Year())

	birthDateIndoString := fmt.Sprintf("%02d", nurse.BirthDate.Day()) + " " +
		constants.Bulan[int(nurse.BirthDate.Month())] + " " +
		strconv.Itoa(nurse.BirthDate.Year())

	speciality, err := uc.spcRep.GetById(nurse.SpecialityId)
	if err != nil {
		return res, err
	}

	res = dto.NurseRes{
		ID:                  nurse.ID,
		UserID:              nurse.UserId,
		DoctorID:            nurse.DoctorId,
		SpecialityId:        nurse.SpecialityId,
		LicenseNumber:       nurse.LicenseNumber,
		SpecialityName:      speciality.Name,
		BirthDate:           nurse.BirthDate,
		BirthDateString:     birthDateString,
		BirthDateStringIndo: birthDateIndoString,
		ProfilePic:          nurse.ProfilePic,
		Phone:               nurse.Phone,
		MaritalStatus:       nurse.MaritalStatus,
		Email:               nurse.Email,
		Nira:                nurse.Nira,
		SIP:                 nurse.SIP,
		CreatedAt:           nurse.CreatedAt,
		UpdatedAt:           nurse.UpdatedAt,
		DeletedAt:           nurse.DeletedAt,
	}

	return res, nil
}

func (uc *nurseUseCase) Delete(id uint) error {

	err := uc.nurseRep.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
