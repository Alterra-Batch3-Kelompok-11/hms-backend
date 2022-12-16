package mock

import (
	"errors"
	"hms-backend/dto"

	"github.com/stretchr/testify/mock"
)

type ReligionUseCaseMock struct {
	mock.Mock
}

func New() *ReligionUseCaseMock {
	return &ReligionUseCaseMock{}
}

func (uc *ReligionUseCaseMock) GetAll() ([]dto.Religion, error) {
	args := uc.Called()

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).([]dto.Religion)

	if isSuccess {
		return data, nil
	} else {
		return []dto.Religion{}, errors.New("testing error")
	}
}

func (uc *ReligionUseCaseMock) GetById(id uint) (dto.Religion, error) {
	args := uc.Called(id)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(dto.Religion)

	if isSuccess {
		return data, nil
	} else {
		return dto.Religion{}, errors.New("testing error")
	}
}
