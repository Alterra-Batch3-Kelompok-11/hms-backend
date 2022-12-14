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
	"time"
)

type AuthUseCase interface {
	Login(username, password string) (dto.LoginRes, error)
	SignUp(payload dto.UserReq) (dto.UserRes, error)
	RefreshToken(id uint) (dto.LoginRes, error)
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
		return dto.LoginRes{}, errors.New("username or password is wrong")
	}

	checkPass, err := helpers.CheckPasswordHash(password, user.Password)
	if err != nil {
		return dto.LoginRes{}, errors.New("username or password is wrong")
	}
	if !checkPass {
		return dto.LoginRes{}, errors.New("username or password is wrong")
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
		Token:    token,
		DoctorID: nil,
		NurseID:  nil,
	}

	if user.RoleId == 2 {
		doctor, err := uc.doctorRepository.GetByUserId(user.ID)
		if err != nil {
			return dto.LoginRes{}, err
		}

		res.DoctorID = doctor.ID
	} else if user.RoleId == 3 {
		nurse, err := uc.nurseRepository.GetByUserId(user.ID)
		if err != nil {
			return dto.LoginRes{}, err
		}

		res.NurseID = nurse.ID
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
		existDoctor, _ := uc.doctorRepository.GetByLicenseNumber(payload.LicenseNumber)
		if existDoctor.ID != 0 {
			return dto.UserRes{}, errors.New("license number already exist")
		}
	} else if payload.RoleID == 3 {
		existNurse, _ := uc.nurseRepository.GetByLicenseNumber(payload.LicenseNumber)
		if existNurse.ID != 0 {
			return dto.UserRes{}, errors.New("license number already exist")
		}
	}

	// TODO check username exist
	usernameExist, _ := uc.userRepository.GetByUsername(payload.Username)
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
			ProfilePic:    payload.ProfilePic,
			BirthDate:     time.Time{},
			Phone:         "",
			MaritalStatus: false,
			Email:         "",
			User:          models.User{},
			Speciality:    models.Speciality{},
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
func (uc *authUseCase) RefreshToken(id uint) (dto.LoginRes, error) {

	user, err := uc.userRepository.GetById(id)
	if err != nil {
		return dto.LoginRes{}, errors.New("user not found")
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
		Token:    token,
		DoctorID: nil,
		NurseID:  nil,
	}

	if user.RoleId == 2 {
		doctor, err := uc.doctorRepository.GetByUserId(user.ID)
		if err != nil {
			return dto.LoginRes{}, err
		}

		res.DoctorID = doctor.ID
	} else if user.RoleId == 3 {
		nurse, err := uc.nurseRepository.GetByUserId(user.ID)
		if err != nil {
			return dto.LoginRes{}, err
		}

		res.NurseID = nurse.ID
	}

	return res, nil
}
