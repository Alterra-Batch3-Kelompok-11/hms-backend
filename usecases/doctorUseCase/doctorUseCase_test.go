package doctorUseCase

import (
	"hms-backend/constants"
	"hms-backend/dto"
	"hms-backend/models"
	"hms-backend/repositories/doctorRepository/mock"
	mock4 "hms-backend/repositories/doctorScheduleRepository/mock"
	mock3 "hms-backend/repositories/specialityRepository/mock"
	mock2 "hms-backend/repositories/userRepository/mock"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type doctorTestSuite struct {
	suite.Suite
	usecase   *doctorUseCase
	doctorRep *mock.DoctorRepositoryMock
	userRep   *mock2.UserRepositoryMock
	spcRep    *mock3.SpecialityRepositoryMock
	scdRep    *mock4.DoctorScheduleRepositoryMock
}

func (s *doctorTestSuite) SetupSuite() {
	s.doctorRep = &mock.DoctorRepositoryMock{}
	s.userRep = &mock2.UserRepositoryMock{}
	s.spcRep = &mock3.SpecialityRepositoryMock{}
	s.scdRep = &mock4.DoctorScheduleRepositoryMock{}

	s.usecase = &doctorUseCase{
		doctorRep: s.doctorRep,
		userRep:   s.userRep,
		spcRep:    s.spcRep,
		scdRep:    s.scdRep,
	}

	s.doctorRep.On("GetAll").Return(
		true,
		Doctors,
	).On("GetById", uint(1)).Return(
		true,
		Doctors[0],
	).On("GetById", uint(2)).Return(
		false,
		models.Doctor{},
	).On("GetByLicenseNumber", "1234567890").Return(
		true,
		Doctors[0],
	).On("GetByLicenseNumber", "1234567891").Return(
		false,
		models.Doctor{},
	).On("GetBySpecialityId", uint(1)).Return(
		true,
		Doctors,
	).On("GetBySpecialityId", uint(0)).Return(
		false,
		[]models.Doctor{},
	).On("Create", constants.Anything).Return(
		true,
		Doctors[0],
	).On("GetByLicenseNumberOther", "1234567890", uint(1)).Return(
		true,
		Doctors[0],
	).On("GetByLicenseNumberOther", "1234567891", uint(1)).Return(
		false,
		models.Doctor{},
	).On("Update", uint(1), constants.Anything).Return(
		true,
		Doctors[0],
	).On("Delete", uint(1)).Return(
		true,
	).On("Delete", uint(2)).Return(
		false,
	)

	s.userRep.On("GetById", uint(1)).Return(
		true,
		Users[0],
	).On("GetByUsername", "1234567890").Return(
		true,
		Users[0],
	).On("GetByUsername", "1234567891").Return(
		false,
		models.User{},
	).On("Create", constants.Anything).Return(
		true,
		Users[0],
	).On("Update", uint(1), constants.Anything).Return(
		true,
		Users[0],
	).On("Delete", uint(1)).Return(
		true,
	).On("Delete", uint(2)).Return(
		false,
	)

	s.spcRep.On("GetById", uint(1)).Return(
		true,
		Speciality,
	)

	s.scdRep.On("GetByDoctorIdDay", uint(1), 0).Return(
		false,
		models.DoctorSchedule{},
	).On("GetByDoctorIdDay", uint(1), 1).Return(
		true,
		Schedules[0],
	).On("GetByDoctorIdDay", uint(1), 2).Return(
		false,
		models.DoctorSchedule{},
	).On("GetByDoctorIdDay", uint(1), 3).Return(
		false,
		models.DoctorSchedule{},
	).On("GetByDoctorIdDay", uint(1), 4).Return(
		false,
		models.DoctorSchedule{},
	).On("GetByDoctorIdDay", uint(1), 5).Return(
		false,
		models.DoctorSchedule{},
	).On("GetByDoctorIdDay", uint(1), 6).Return(
		false,
		models.DoctorSchedule{},
	).On("GetByDoctorIdDay", uint(0), 0).Return(
		false,
		models.DoctorSchedule{},
	).On("GetByDoctorIdDay", uint(0), 1).Return(
		true,
		Schedules[0],
	).On("GetByDoctorIdDay", uint(0), 2).Return(
		false,
		models.DoctorSchedule{},
	).On("GetByDoctorIdDay", uint(0), 3).Return(
		false,
		models.DoctorSchedule{},
	).On("GetByDoctorIdDay", uint(0), 4).Return(
		false,
		models.DoctorSchedule{},
	).On("GetByDoctorIdDay", uint(0), 5).Return(
		false,
		models.DoctorSchedule{},
	).On("GetByDoctorIdDay", uint(0), 6).Return(
		false,
		models.DoctorSchedule{},
	).On("GetByDay", 0).Return(
		true,
		Schedules,
	).On("GetByDay", 1).Return(
		true,
		Schedules,
	).On("GetByDay", 2).Return(
		true,
		Schedules,
	).On("GetByDay", 3).Return(
		true,
		Schedules,
	).On("GetByDay", 4).Return(
		true,
		Schedules,
	).On("GetByDay", 5).Return(
		true,
		Schedules,
	).On("GetByDay", 6).Return(
		true,
		Schedules,
	)
}

func TestOrderSuite(t *testing.T) {
	suite.Run(t, new(doctorTestSuite))
}

func (s *doctorTestSuite) TestGetAll() {

	testCases := []struct {
		Name     string
		Expected []dto.DoctorRes
	}{
		{
			"Success",
			ExpectedDoctors,
		},
	}

	for _, testCase := range testCases {
		s.T().Run(testCase.Name, func(t *testing.T) {
			resUc, err := s.usecase.GetAll()

			s.NoError(err)
			for _, res := range resUc {
				s.Equal(testCase.Expected[0].ID, res.ID)
				s.Equal(testCase.Expected[0].Name, res.Name)
				s.Equal(testCase.Expected[0].Email, res.Email)
				s.Equal(testCase.Expected[0].BirthDate, res.BirthDate)
				s.Equal(testCase.Expected[0].LicenseNumber, res.LicenseNumber)
				s.Equal(testCase.Expected[0].SpecialityName, res.SpecialityName)
			}
		})
	}
}
func (s *doctorTestSuite) TestGetById() {
	testCases := []struct {
		Name      string
		IsSuccess bool
		Id        uint
		Expected  dto.DoctorRes
	}{
		{
			"Success",
			true,
			1,
			ExpectedDoctors[0],
		},
		{
			"Failed",
			false,
			2,
			ExpectedDoctors[0],
		},
	}

	for _, testCase := range testCases {
		s.T().Run(testCase.Name, func(t *testing.T) {
			resUc, err := s.usecase.GetById(testCase.Id)

			if testCase.IsSuccess {
				s.NoError(err)
				s.Equal(testCase.Expected.ID, resUc.ID)
				s.Equal(testCase.Expected.Name, resUc.Name)
				s.Equal(testCase.Expected.Email, resUc.Email)
				s.Equal(testCase.Expected.BirthDate, resUc.BirthDate)
				s.Equal(testCase.Expected.LicenseNumber, resUc.LicenseNumber)
				s.Equal(testCase.Expected.SpecialityName, resUc.SpecialityName)
			} else {
				s.Equal("testing error", err.Error())
			}
		})
	}
}
func (s *doctorTestSuite) TestGetByLicenseNumber() {
	testCases := []struct {
		Name          string
		IsSuccess     bool
		LicenseNumber string
		Expected      dto.DoctorRes
	}{
		{
			"Success",
			true,
			"1234567890",
			ExpectedDoctors[0],
		},
		{
			"Failed",
			false,
			"1234567891",
			ExpectedDoctors[0],
		},
	}

	for _, testCase := range testCases {
		s.T().Run(testCase.Name, func(t *testing.T) {
			resUc, err := s.usecase.GetByLicenseNumber(testCase.LicenseNumber)

			if testCase.IsSuccess {
				s.NoError(err)
				s.Equal(testCase.Expected.ID, resUc.ID)
				s.Equal(testCase.Expected.Name, resUc.Name)
				s.Equal(testCase.Expected.Email, resUc.Email)
				s.Equal(testCase.Expected.BirthDate, resUc.BirthDate)
				s.Equal(testCase.Expected.LicenseNumber, resUc.LicenseNumber)
				s.Equal(testCase.Expected.SpecialityName, resUc.SpecialityName)
			} else {
				s.Equal("testing error", err.Error())
			}
		})
	}
}
func (s *doctorTestSuite) TestGetBySpecialityId() {
	testCases := []struct {
		Name         string
		IsSuccess    bool
		SpecialityId uint
		Expected     []dto.DoctorRes
	}{
		{
			"Success",
			true,
			1,
			ExpectedDoctors,
		},
		{
			"Failed",
			false,
			0,
			nil,
		},
	}

	for _, testCase := range testCases {
		s.T().Run(testCase.Name, func(t *testing.T) {
			resUc, err := s.usecase.GetBySpecialityId(testCase.SpecialityId)

			if testCase.IsSuccess {
				s.NoError(err)
				for _, res := range resUc {
					s.Equal(testCase.Expected[0].ID, res.ID)
					s.Equal(testCase.Expected[0].Name, res.Name)
					s.Equal(testCase.Expected[0].Email, res.Email)
					s.Equal(testCase.Expected[0].BirthDate, res.BirthDate)
					s.Equal(testCase.Expected[0].LicenseNumber, res.LicenseNumber)
					s.Equal(testCase.Expected[0].SpecialityName, res.SpecialityName)
				}
			} else {
				s.Equal("testing error", err.Error())
			}

		})
	}
}
func (s *doctorTestSuite) TestGetToday() {
	testCases := []struct {
		Name     string
		Expected []dto.DoctorRes
	}{
		{
			"Success",
			ExpectedDoctors,
		},
	}

	for _, testCase := range testCases {
		s.T().Run(testCase.Name, func(t *testing.T) {
			resUc, err := s.usecase.GetToday()

			s.NoError(err)
			for _, res := range resUc {
				s.Equal(testCase.Expected[0].ID, res.ID)
				s.Equal(testCase.Expected[0].Name, res.Name)
				s.Equal(testCase.Expected[0].Email, res.Email)
				s.Equal(testCase.Expected[0].BirthDate, res.BirthDate)
				s.Equal(testCase.Expected[0].LicenseNumber, res.LicenseNumber)
				s.Equal(testCase.Expected[0].SpecialityName, res.SpecialityName)
			}
		})
	}
}
func (s *doctorTestSuite) TestCreate() {
	testCases := []struct {
		Name      string
		IsSuccess bool
		Payload   dto.DoctorReq
		Expected  dto.DoctorRes
	}{
		{
			"Success",
			true,
			PayloadDoctor,
			ExpectedDoctors[0],
		},
		{
			"Failed",
			false,
			dto.DoctorReq{
				Name:          PayloadDoctor.Name,
				LicenseNumber: "1234567890",
				Password:      "",
				SpecialityID:  0,
				ProfilePic:    "",
				BirthDate:     "",
				Phone:         "",
				MaritalStatus: false,
				Email:         "",
			},
			dto.DoctorRes{},
		},
	}

	for _, testCase := range testCases {
		s.T().Run(testCase.Name, func(t *testing.T) {
			resUc, err := s.usecase.Create(testCase.Payload)

			if testCase.IsSuccess {
				s.NoError(err)
				s.Equal(testCase.Expected.ID, resUc.ID)
				s.Equal(testCase.Expected.Name, resUc.Name)
				s.Equal(testCase.Expected.Email, resUc.Email)
				s.Equal(testCase.Expected.SpecialityName, resUc.SpecialityName)
			} else {
				s.Equal("license number already exist", err.Error())
			}

		})
	}
}
func (s *doctorTestSuite) TestUpdate() {
	testCases := []struct {
		Name      string
		IsSuccess bool
		Id        uint
		Payload   dto.DoctorReq
		Expected  dto.DoctorRes
	}{
		{
			"Success",
			true,
			uint(1),
			PayloadDoctor,
			ExpectedDoctors[0],
		},
		{
			"Failed",
			false,
			uint(2),
			dto.DoctorReq{
				Name:          PayloadDoctor.Name,
				LicenseNumber: "1234567891",
				Password:      "",
				SpecialityID:  0,
				ProfilePic:    "",
				BirthDate:     "",
				Phone:         "",
				MaritalStatus: false,
				Email:         "",
			},
			dto.DoctorRes{},
		},
	}

	for _, testCase := range testCases {
		s.T().Run(testCase.Name, func(t *testing.T) {
			resUc, err := s.usecase.Update(testCase.Id, testCase.Payload)

			if testCase.IsSuccess {
				s.NoError(err)
				s.Equal(testCase.Expected.ID, resUc.ID)
				s.Equal(testCase.Expected.Name, resUc.Name)
				s.Equal(testCase.Expected.Email, resUc.Email)
				s.Equal(testCase.Expected.SpecialityName, resUc.SpecialityName)
			} else {
				s.Equal("testing error", err.Error())
			}

		})
	}
}
func (s *doctorTestSuite) TestDelete() {
	testCases := []struct {
		Name      string
		IsSuccess bool
		Id        uint
		Expected  string
	}{
		{
			"Success",
			true,
			uint(1),
			"",
		},
		{
			"Failed",
			false,
			uint(2),
			"testing error",
		},
	}

	for _, testCase := range testCases {
		s.T().Run(testCase.Name, func(t *testing.T) {
			err := s.usecase.Delete(testCase.Id)

			if testCase.IsSuccess {
				s.NoError(err)
			} else {
				s.Equal(testCase.Expected, err.Error())
			}

		})
	}
}

var PayloadDoctor = dto.DoctorReq{
	Name:          "Fulan",
	LicenseNumber: "1234567891",
	Password:      "",
	SpecialityID:  1,
	ProfilePic:    "",
	BirthDate:     "1998-01-02",
	Phone:         "08123456789",
	MaritalStatus: false,
	Email:         "fulan@mail.com",
}

var ExpectedDoctors = []dto.DoctorRes{
	{
		ID:                  Doctors[0].ID,
		CreatedAt:           Doctors[0].CreatedAt,
		UpdatedAt:           Doctors[0].UpdatedAt,
		DeletedAt:           Doctors[0].DeletedAt,
		Name:                Users[0].Name,
		SpecialityId:        Doctors[0].SpecialityId,
		LicenseNumber:       Doctors[0].LicenseNumber,
		SpecialityName:      Speciality.Name,
		ProfilePic:          Doctors[0].ProfilePic,
		BirthDate:           Doctors[0].BirthDate,
		BirthDateString:     "",
		BirthDateStringIndo: "",
		Phone:               Doctors[0].Phone,
		MaritalStatus:       Doctors[0].MaritalStatus,
		Email:               Doctors[0].Email,
	},
}

var Doctors = []models.Doctor{
	{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		UserId:        1,
		SpecialityId:  1,
		LicenseNumber: "1234567890",
		ProfilePic:    "",
		BirthDate:     time.Time{},
		Phone:         "081234567890",
		MaritalStatus: false,
		Email:         "fulan@mail.com",
	},
}

var Users = []models.User{
	{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		RoleId:   2,
		Username: "1234567890",
		Password: "ini_password",
		Name:     "Dr. Fulan",
	},
}

var Speciality = models.Speciality{
	Model: gorm.Model{
		ID:        1,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: gorm.DeletedAt{},
	},
	Name: "Umum",
}

var Schedules = []models.DoctorSchedule{
	{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		DoctorId:  1,
		Day:       1,
		StartTime: "08:00",
		EndTime:   "16:00",
	},
}
