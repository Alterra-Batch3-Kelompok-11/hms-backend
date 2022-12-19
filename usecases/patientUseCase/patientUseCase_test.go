package patientUseCase

import (
	"hms-backend/dto"
	"hms-backend/models"
	"hms-backend/repositories/patientRepository/mock"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type patientTestSuite struct {
	suite.Suite
	usecase    *patientUseCase
	patientRep *mock.PatientRepositoryMock
}

func (s *patientTestSuite) SetupSuite() {
	s.patientRep = &mock.PatientRepositoryMock{}

	s.usecase = &patientUseCase{
		patientRep: s.patientRep,
	}

	s.patientRep.On("GetAll").Return(
		true,
		Patients,
	).On("GetById", uint(1)).Return(
		true,
		Patients[0],
	).On("GetById", uint(2)).Return(
		false,
		models.Patient{},
	).On("Create", PayloadPatients).Return(
		true,
		Patients[0],
	).On("Create", uint(1), PayloadPatients).Return(
		false,
		models.Patient{},
	).On("Update", uint(1), PayloadPatients).Return(
		true,
		Patients[0],
	).On("Update", uint(2), PayloadPatients).Return(
		false,
		models.Patient{},
	).On("Delete", uint(1)).Return(
		true,
	).On("Delete", uint(2)).Return(
		false,
	)
}

func TestOrderSuite(t *testing.T) {
	suite.Run(t, new(patientTestSuite))
}

func (s *patientTestSuite) TestGetAll() {

	testCases := []struct {
		Name     string
		Expected []dto.PatientRes
	}{
		{
			"Success",
			ExpectedPatients,
		},
	}

	for _, testCase := range testCases {
		s.T().Run(testCase.Name, func(t *testing.T) {
			resUc, err := s.usecase.GetAll()

			s.NoError(err)
			for _, res := range resUc {
				s.Equal(testCase.Expected[0].ID, res.ID)
				s.Equal(testCase.Expected[0].Nik, res.Nik)
				s.Equal(testCase.Expected[0].Name, res.Name)
				s.Equal(testCase.Expected[0].BirthDate, res.BirthDate)
				s.Equal(testCase.Expected[0].Gender, res.Gender)
				s.Equal(testCase.Expected[0].Address, res.Address)
				s.Equal(testCase.Expected[0].Phone, res.Phone)
			}
		})
	}
}

func (s *patientTestSuite) TestGetById() {
	testCases := []struct {
		Name      string
		IsSuccess bool
		Id        uint
		Expected  dto.PatientRes
	}{
		{
			"Success",
			true,
			1,
			ExpectedPatients[0],
		},
		{
			"Failed",
			false,
			2,
			ExpectedPatients[0],
		},
	}

	for _, testCase := range testCases {
		s.T().Run(testCase.Name, func(t *testing.T) {
			resUc, err := s.usecase.GetById(testCase.Id)

			if testCase.IsSuccess {
				s.NoError(err)
				s.Equal(testCase.Expected.ID, resUc.ID)
				s.Equal(testCase.Expected.Nik, resUc.Nik)
				s.Equal(testCase.Expected.Name, resUc.Name)
				s.Equal(testCase.Expected.BirthDate, resUc.BirthDate)
				s.Equal(testCase.Expected.Gender, resUc.Gender)
				s.Equal(testCase.Expected.Address, resUc.Address)
				s.Equal(testCase.Expected.Phone, resUc.Phone)
			} else {
				s.Equal("testing error", err.Error())
			}
		})
	}
}

func (s *patientTestSuite) TestCreate() {
	testCases := []struct {
		Name      string
		IsSuccess bool
		Payload   dto.Patient
		Expected  dto.PatientRes
	}{
		{
			"Success",
			true,
			PayloadPatients,
			ExpectedPatients[0],
		},
		{
			"Failed",
			false,
			dto.Patient{
				NIK:           "12312312123",
				Name:          PayloadPatients.Name,
				BirthDate:     "1997-01-02",
				Gender:        1,
				Phone:         "",
				Address:       "",
				MaritalStatus: false,
				ReligionID:    1,
			},
			dto.PatientRes{},
		},
	}

	for _, testCase := range testCases {
		s.T().Run(testCase.Name, func(t *testing.T) {
			resUc, err := s.usecase.Create(testCase.Payload)

			if testCase.IsSuccess {
				s.NoError(err)
				s.Equal(testCase.Expected.ID, resUc.ID)
				s.Equal(testCase.Expected.Nik, resUc.Nik)
				s.Equal(testCase.Expected.Name, resUc.Name)
				s.Equal(testCase.Expected.BirthDate, resUc.BirthDate)
				s.Equal(testCase.Expected.Gender, resUc.Gender)
				s.Equal(testCase.Expected.Address, resUc.Address)
				s.Equal(testCase.Expected.Phone, resUc.Phone)
				s.Equal(testCase.Expected.MaritalStatus, resUc.MaritalStatus)
				s.Equal(testCase.Expected.ReligionID, resUc.ReligionID)
			} else {
				s.Equal("Nik already exist", err.Error())
			}

		})
	}
}

func (s *patientTestSuite) TestUpdate() {
	testCases := []struct {
		Name      string
		IsSuccess bool
		Id        uint
		Payload   dto.Patient
		Expected  dto.PatientRes
	}{
		{
			"Success",
			true,
			uint(1),
			PayloadPatients,
			ExpectedPatients[10],
		},
		{
			"Failed",
			false,
			uint(2),
			dto.Patient{
				NIK:           "123123123",
				Name:          PayloadPatients.Name,
				BirthDate:     "1997-01-02",
				Gender:        1,
				Phone:         "",
				Address:       "",
				MaritalStatus: false,
				ReligionID:    1,
			},
			dto.PatientRes{},
		},
	}

	for _, testCase := range testCases {
		s.T().Run(testCase.Name, func(t *testing.T) {
			resUc, err := s.usecase.Update(testCase.Id, testCase.Payload)

			if testCase.IsSuccess {
				s.NoError(err)
				s.Equal(testCase.Expected.ID, resUc.ID)
				s.Equal(testCase.Expected.Nik, resUc.Nik)
				s.Equal(testCase.Expected.Name, resUc.Name)
				s.Equal(testCase.Expected.BirthDate, resUc.BirthDate)
				s.Equal(testCase.Expected.Gender, resUc.Gender)
				s.Equal(testCase.Expected.Address, resUc.Address)
				s.Equal(testCase.Expected.Phone, resUc.Phone)
				s.Equal(testCase.Expected.MaritalStatus, resUc.MaritalStatus)
				s.Equal(testCase.Expected.ReligionID, resUc.ReligionID)
			} else {
				s.Equal("testing error", err.Error())
			}

		})
	}
}

func (s *patientTestSuite) TestDelete() {
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

// Data Mock

var PayloadPatients = dto.Patient{
	NIK:           "123123123123",
	Name:          "Test Patient",
	BirthDate:     "1997-01-02",
	Gender:        1,
	Phone:         "081234567890",
	Address:       "Jl. Test",
	MaritalStatus: false,
	ReligionID:    1,
}

var ExpectedPatients = []dto.PatientRes{
	{
		ID:            Patients[0].ID,
		CreatedAt:     Patients[0].CreatedAt,
		UpdatedAt:     Patients[0].UpdatedAt,
		DeletedAt:     Patients[0].DeletedAt,
		Nik:           Patients[0].Nik,
		Name:          Patients[0].Name,
		BirthDate:     Patients[0].BirthDate,
		Gender:        Patients[0].Gender,
		Address:       Patients[0].Address,
		Phone:         Patients[0].Phone,
		MaritalStatus: Patients[0].MaritalStatus,
		ReligionID:    Patients[0].ReligionID,
	},
}

var Patients = []models.Patient{
	{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		ReligionID:    1,
		Nik:           "1234567890123456",
		Name:          "Test Patient",
		BirthDate:     time.Time{},
		Gender:        1,
		Address:       "Jl. Test",
		Phone:         "081234567890",
		MaritalStatus: false,
	},
}
