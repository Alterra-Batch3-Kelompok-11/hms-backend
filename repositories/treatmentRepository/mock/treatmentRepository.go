package mock

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"hms-backend/models"
)

type TreatmentRepositoryMock struct {
	mock.Mock
}

func New() *TreatmentRepositoryMock {
	return &TreatmentRepositoryMock{}
}

func (rep *TreatmentRepositoryMock) GetAll() ([]models.Treatment, error) {
	args := rep.Called()

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).([]models.Treatment)

	if isSuccess {
		return data, nil
	} else {
		return []models.Treatment{}, errors.New("testing error")
	}
}
func (rep *TreatmentRepositoryMock) GetById(id uint) (models.Treatment, error) {
	args := rep.Called(id)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(models.Treatment)

	if isSuccess {
		return data, nil
	} else {
		return models.Treatment{}, errors.New("testing error")
	}
}
func (rep *TreatmentRepositoryMock) GetByOutpatientSessionId(id uint) (models.Treatment, error) {
	args := rep.Called(id)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(models.Treatment)

	if isSuccess {
		return data, nil
	} else {
		return models.Treatment{}, errors.New("testing error")
	}
}
func (rep *TreatmentRepositoryMock) Create(payload models.Treatment) (models.Treatment, error) {
	args := rep.Called(payload)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(models.Treatment)

	if isSuccess {
		return data, nil
	} else {
		return models.Treatment{}, errors.New("testing error")
	}
}
func (rep *TreatmentRepositoryMock) Update(id uint, payload models.Treatment) (models.Treatment, error) {
	args := rep.Called(id, payload)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(models.Treatment)

	if isSuccess {
		return data, nil
	} else {
		return models.Treatment{}, errors.New("testing error")
	}
}
func (rep *TreatmentRepositoryMock) Delete(id uint) error {
	args := rep.Called(id)

	isSuccess := args.Get(0).(bool)

	if isSuccess {
		return nil
	} else {
		return errors.New("testing error")
	}
}
