package religionUseCase

import (
	"hms-backend/dto"
	"hms-backend/models"
	"hms-backend/repositories/religionRepository/mock"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type religionTestSuite struct {
	suite.Suite
	usecase     *religionUseCase
	religionRep *mock.ReligionRepositoryMock
}

func (s *religionTestSuite) SetupSuite() {
	s.religionRep = &mock.ReligionRepositoryMock{}

	s.usecase = &religionUseCase{
		s.religionRep,
	}

	s.religionRep.On("GetAll").Return(
		true,
		Religions,
	).On("GetById", uint(1)).Return(
		true,
		Religions[0],
	).On("GetById", uint(2)).Return(
		false,
		models.Religion{},
	)
}

func TestOrderSuite(t *testing.T) {
	suite.Run(t, new(religionTestSuite))
}

func (s *religionTestSuite) TestGetAll() {

	testCases := []struct {
		Name     string
		Expected []dto.Religion
	}{
		{
			"Success",
			ExpectedReligions,
		},
	}

	for _, testCase := range testCases {
		s.T().Run(testCase.Name, func(t *testing.T) {
			resUc, err := s.usecase.GetAll()

			s.NoError(err)
			for _, res := range resUc {
				s.Equal(testCase.Expected[0].ID, res.ID)
				s.Equal(testCase.Expected[0].Name, res.Name)
			}
		})
	}
}

func (s *religionTestSuite) TestGetById() {
	testCases := []struct {
		Name      string
		IsSuccess bool
		Id        uint
		Expected  dto.Religion
	}{
		{
			"Success",
			true,
			1,
			ExpectedReligions[0],
		},
		{
			"Failed",
			false,
			2,
			ExpectedReligions[0],
		},
	}

	for _, testCase := range testCases {
		s.T().Run(testCase.Name, func(t *testing.T) {
			resUc, err := s.usecase.GetById(testCase.Id)

			if testCase.IsSuccess {
				s.NoError(err)
				s.Equal(testCase.Expected.ID, resUc.ID)
				s.Equal(testCase.Expected.Name, resUc.Name)
			} else {
				s.Equal("testing error", err.Error())
			}
		})
	}
}

// Data Mock

var PayloadReligion = dto.Religion{
	ID:        1,
	CreatedAt: time.Time{},
	UpdatedAt: time.Time{},
	DeletedAt: gorm.DeletedAt{},
	Name:      "Islam",
}

var ExpectedReligions = []dto.Religion{
	{
		ID:        Religions[0].ID,
		CreatedAt: Religions[0].CreatedAt,
		UpdatedAt: Religions[0].UpdatedAt,
		DeletedAt: Religions[0].DeletedAt,
		Name:      Religions[0].Name,
	},
}

var Religions = []models.Religion{
	{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		Name: "Admin",
	},
}
