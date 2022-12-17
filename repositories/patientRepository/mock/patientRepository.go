package mock

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"hms-backend/models"
)

type PatientRepositoryMock struct {
	mock.Mock
}

func New() *PatientRepositoryMock {
	return &PatientRepositoryMock{}
}

func (rep *PatientRepositoryMock) GetAll() ([]models.Patient, error) {
	args := rep.Called()

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).([]models.Patient)

	if isSuccess {
		return data, nil
	} else {
		return []models.Patient{}, errors.New("testing error")
	}
}
func (rep *PatientRepositoryMock) GetById(id uint) (models.Patient, error) {
	args := rep.Called(id)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(models.Patient)

	if isSuccess {
		return data, nil
	} else {
		return models.Patient{}, errors.New("testing error")
	}
}
func (rep *PatientRepositoryMock) Count() (int64, error) {
	args := rep.Called()

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(int64)

	if isSuccess {
		return data, nil
	} else {
		return 0, errors.New("testing error")
	}
}
func (rep *PatientRepositoryMock) Create(payload models.Patient) (models.Patient, error) {
	args := rep.Called(payload)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(models.Patient)

	if isSuccess {
		return data, nil
	} else {
		return models.Patient{}, errors.New("testing error")
	}
}
func (rep *PatientRepositoryMock) Update(id uint, payload models.Patient) (models.Patient, error) {
	args := rep.Called(id, payload)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(models.Patient)

	if isSuccess {
		return data, nil
	} else {
		return models.Patient{}, errors.New("testing error")
	}
}
func (rep *PatientRepositoryMock) Delete(id uint) error {
	args := rep.Called(id)

	isSuccess := args.Get(0).(bool)

	if isSuccess {
		return nil
	} else {
		return errors.New("testing error")
	}
}
