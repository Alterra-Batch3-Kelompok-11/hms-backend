package authUseCase

import (
	"errors"
	"gorm.io/gorm"
	"hms-backend/dto"
	"hms-backend/helpers"
	"hms-backend/middlewares"
	"hms-backend/models"
	"hms-backend/repositories/doctorRepository"
	"hms-backend/repositories/nurseRepository"
	"hms-backend/repositories/userRepository"
)

type AuthUseCase interface {
	Login(username, password string) (dto.LoginRes, error)
	SignUp(payload dto.UserReq) (dto.UserRes, error)
}

type authUseCase struct {
	userRepository   userRepository.UserRepository
	doctorRepository doctorRepository.DoctorRepository
	nurseRepository  nurseRepository.NurseRepository
}

func New(rep userRepository.UserRepository, dtrRep doctorRepository.DoctorRepository, nrsRep nurseRepository.NurseRepository) *authUseCase {
	return &authUseCase{rep, dtrRep, nrsRep}
}

func (uc *authUseCase) Login(username, password string) (dto.LoginRes, error) {

	user, err := uc.userRepository.GetByUsername(username)
	if err != nil {
		return dto.LoginRes{}, err
	}

	checkPass, err := helpers.CheckPasswordHash(password, user.Password)
	if err != nil {
		return dto.LoginRes{}, err
	}
	if !checkPass {
		return dto.LoginRes{}, errors.New("record not found")
	}

	token, err := middlewares.CreateToken(user.ID, user.Username, user.RoleId)
	if err != nil {
		return dto.LoginRes{}, err
	}

	res := dto.LoginRes{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
		RoleID:   user.RoleId,
		Role:     user.Role.Name,
		Token:    token,
	}

	return res, nil
}

func (uc *authUseCase) SignUp(payload dto.UserReq) (dto.UserRes, error) {

	hashedPass, err := helpers.HashPassword(payload.Password)
	if err != nil {
		return dto.UserRes{}, err
	}

	// TODO Check License Number if Exist
	if payload.RoleID == 2 {
		existDoctor, err := uc.doctorRepository.GetByLicenseNumber(payload.LicenseNumber)
		if err != nil {
			return dto.UserRes{}, err
		}
		if existDoctor.ID != 0 {
			return dto.UserRes{}, errors.New("license number already exist")
		}
	} else if payload.RoleID == 3 {
		existNurse, err := uc.nurseRepository.GetByLicenseNumber(payload.LicenseNumber)
		if err != nil {
			return dto.UserRes{}, err
		}
		if existNurse.ID != 0 {
			return dto.UserRes{}, errors.New("license number already exist")
		}
	}

	// TODO check username exist
	usernameExist, err := uc.userRepository.GetByUsername(payload.Username)
	if err != nil {
		return dto.UserRes{}, err
	}
	if usernameExist.ID != 0 {
		return dto.UserRes{}, errors.New("username already taken")
	}

	user := models.User{
		Model:    gorm.Model{},
		RoleId:   payload.RoleID,
		Username: payload.Username,
		Password: hashedPass,
		Name:     payload.Name,
	}

	resCreateUsr, err := uc.userRepository.Create(user)
	if err != nil {
		return dto.UserRes{}, err
	}

	licenseNumber := ""

	if payload.RoleID == 2 { // if role as doctor
		doctor := models.Doctor{
			Model:         gorm.Model{},
			UserId:        resCreateUsr.ID,
			SpecialityId:  payload.SpecialityID,
			LicenseNumber: payload.LicenseNumber,
		}

		resCreateDtr, err := uc.doctorRepository.Create(doctor)
		if err != nil {
			return dto.UserRes{}, err
		}

		licenseNumber = resCreateDtr.LicenseNumber
	} else if payload.RoleID == 3 { // if role as nurse
		nurse := models.Nurse{
			Model:         gorm.Model{},
			UserId:        resCreateUsr.ID,
			LicenseNumber: payload.LicenseNumber,
		}

		resCreateNrs, err := uc.nurseRepository.Create(nurse)
		if err != nil {
			return dto.UserRes{}, err
		}

		licenseNumber = resCreateNrs.LicenseNumber
	}

	res := dto.UserRes{
		ID:            resCreateUsr.ID,
		CreatedAt:     resCreateUsr.CreatedAt,
		UpdatedAt:     resCreateUsr.UpdatedAt,
		DeletedAt:     resCreateUsr.DeletedAt,
		Name:          resCreateUsr.Name,
		Username:      resCreateUsr.Username,
		RoleID:        resCreateUsr.RoleId,
		Role:          resCreateUsr.Role.Name,
		Password:      resCreateUsr.Password,
		LicenseNumber: licenseNumber,
	}

	return res, err
}