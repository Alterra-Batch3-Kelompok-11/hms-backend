package mock

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"hms-backend/models"
)

type ReligionRepositoryMock struct {
	mock.Mock
}

func New() *ReligionRepositoryMock {
	return &ReligionRepositoryMock{}
}

func (rep *ReligionRepositoryMock) GetAll() ([]models.Religion, error) {
	args := rep.Called()

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).([]models.Religion)

	if isSuccess {
		return data, nil
	} else {
		return []models.Religion{}, errors.New("testing error")
	}
}
func (rep *ReligionRepositoryMock) GetById(id uint) (models.Religion, error) {
	args := rep.Called(id)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(models.Religion)

	if isSuccess {
		return data, nil
	} else {
		return models.Religion{}, errors.New("testing error")
	}
}
