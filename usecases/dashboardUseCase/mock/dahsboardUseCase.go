package mock

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"hms-backend/dto"
)

type DashboardUseCaseMock struct {
	mock.Mock
}

func New() *DashboardUseCaseMock {
	return &DashboardUseCaseMock{}
}

func (uc *DashboardUseCaseMock) GetDataDashboardWeb() (dto.DashboardWeb, error) {
	args := uc.Called()

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(dto.DashboardWeb)

	if isSuccess {
		return data, nil
	} else {
		return dto.DashboardWeb{}, errors.New("testing error")
	}
}
func (uc *DashboardUseCaseMock) GetDataDashboardMobile(doctorId uint) (dto.DashboardMobile, error) {
	args := uc.Called(doctorId)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(dto.DashboardMobile)

	if isSuccess {
		return data, nil
	} else {
		return dto.DashboardMobile{}, errors.New("testing error")
	}
}
