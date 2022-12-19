package mock

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"hms-backend/dto"
)

type HistoryUseCaseMock struct {
	mock.Mock
}

func New() *HistoryUseCaseMock {
	return &HistoryUseCaseMock{}
}

func (uc *HistoryUseCaseMock) GetOutpatientSessionHistory(doctorId uint) ([]dto.History, error) {
	args := uc.Called(doctorId)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).([]dto.History)

	if isSuccess {
		return data, nil
	} else {
		return []dto.History{}, errors.New("testing error")
	}
}
func (uc *HistoryUseCaseMock) GetApprovalHistory(doctorId uint) ([]dto.History, error) {
	args := uc.Called(doctorId)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).([]dto.History)

	if isSuccess {
		return data, nil
	} else {
		return []dto.History{}, errors.New("testing error")
	}
}
