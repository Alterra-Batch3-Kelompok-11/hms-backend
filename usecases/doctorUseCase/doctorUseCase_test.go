package doctorUseCase

import (
	"github.com/stretchr/testify/suite"
	"hms-backend/repositories/doctorRepository/mock"
	"testing"
)

type doctorTestSuite struct {
	suite.Suite
	usecase *doctorUseCase
	mock    *mock.DoctorRepositoryMock
}

func (s *doctorTestSuite) SetupSuite() {
	mock := &mock.DoctorRepositoryMock{}
	s.mock = mock

	s.usecase = &doctorUseCase{
		doctorRep: nil,
		userRep:   nil,
		spcRep:    nil,
		scdRep:    nil,
	}
}

func TestOrderSuite(t *testing.T) {
	suite.Run(t, new(doctorTestSuite))
}
