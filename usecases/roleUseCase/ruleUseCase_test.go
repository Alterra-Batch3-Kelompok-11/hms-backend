package roleUseCase

import (
	"hms-backend/dto"
	"hms-backend/models"
	"hms-backend/repositories/roleRepository/mock"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type roleTestSuite struct {
	suite.Suite
	usecase *roleUseCase
	roleRep *mock.RoleRepositoryMock
}

func (s *roleTestSuite) SetupSuite() {
	s.roleRep = &mock.RoleRepositoryMock{}

	s.usecase = &roleUseCase{
		s.roleRep,
	}

	s.roleRep.On("GetAll").Return(
		true,
		Roles,
	).On("GetById", uint(1)).Return(
		true,
		Roles[0],
	).On("GetById", uint(2)).Return(
		false,
		models.Role{},
	)
}

func TestOrderSuite(t *testing.T) {
	suite.Run(t, new(roleTestSuite))
}

func (s *roleTestSuite) TestGetAll() {

	testCases := []struct {
		Name     string
		Expected []dto.Role
	}{
		{
			"Success",
			ExpectedRoles,
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

func (s *roleTestSuite) TestGetById() {
	testCases := []struct {
		Name      string
		IsSuccess bool
		Id        uint
		Expected  dto.Role
	}{
		{
			"Success",
			true,
			1,
			ExpectedRoles[0],
		},
		{
			"Failed",
			false,
			2,
			ExpectedRoles[0],
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

var PayloadRole = dto.Role{
	ID:        1,
	CreatedAt: time.Time{},
	UpdatedAt: time.Time{},
	DeletedAt: gorm.DeletedAt{},
	Name:      "Admin",
}

var ExpectedRoles = []dto.Role{
	{
		ID:        Roles[0].ID,
		CreatedAt: Roles[0].CreatedAt,
		UpdatedAt: Roles[0].UpdatedAt,
		DeletedAt: Roles[0].DeletedAt,
		Name:      Roles[0].Name,
	},
}

var Roles = []models.Role{
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
