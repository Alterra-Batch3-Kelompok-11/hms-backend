package mock

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"hms-backend/dto"
)

type NotificationUseCaseMock struct {
	mock.Mock
}

func New() *NotificationUseCaseMock {
	return &NotificationUseCaseMock{}
}

func (uc *NotificationUseCaseMock) GetByUserId(userId uint) ([]dto.Notification, error) {
	args := uc.Called(userId)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).([]dto.Notification)

	if isSuccess {
		return data, nil
	} else {
		return []dto.Notification{}, errors.New("testing error")
	}
}
