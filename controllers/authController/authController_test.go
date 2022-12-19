package authController

import (
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"hms-backend/dto"
	"hms-backend/middlewares"
	"hms-backend/usecases/authUseCase/mock"
	"testing"
	"time"
)

type dashboardTestSuite struct {
	suite.Suite
	ctrl *authController
	mock *mock.AuthUseCaseMock
}

func (s *dashboardTestSuite) SetupSuite() {
	mock := &mock.AuthUseCaseMock{}
	s.mock = mock

	s.ctrl = &authController{
		usecase: s.mock,
	}
}

func TestAuthSuite(t *testing.T) {
	suite.Run(t, new(dashboardTestSuite))
}

func (s *dashboardTestSuite) Login() {
	s.mock.On("Login").Return(
		true,
		"fulan",
		"fulan123",
	)
}
func (s *dashboardTestSuite) SignUp() {

}
func (s *dashboardTestSuite) RefreshToken() {

}

var Token, _ = middlewares.CreateToken(1, "fulan", 2)

var LoginRes = dto.LoginRes{
	ID:            1,
	Name:          "Fulan",
	Username:      "fulan",
	RoleID:        2,
	Token:         Token,
	LicenseNumber: "1234567890",
	DoctorID:      1,
	NurseID:       nil,
}

var UserRes = dto.UserRes{
	ID:            1,
	CreatedAt:     time.Time{},
	UpdatedAt:     time.Time{},
	DeletedAt:     gorm.DeletedAt{},
	Name:          "Fulan",
	Username:      "fulan",
	RoleID:        2,
	Role:          "Doctor",
	Password:      "oiandfknsodignisrngoetgbeotng",
	LicenseNumber: "1234567890",
}

var UserRe = dto.UserReq{
	Name:          "Fulan",
	LicenseNumber: "1234567890",
	Username:      "fulan",
	Password:      "asdasdasd",
	RoleID:        2,
	SpecialityID:  1,
	ProfilePic:    "",
	BirthDate:     "14-12-1998",
	Phone:         "",
	MaritalStatus: false,
	Email:         "",
	Nira:          "",
	SIP:           "",
}
