package mock

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"hms-backend/models"
)

type HistoryRepositoryMock struct {
	mock.Mock
}

func New() *HistoryRepositoryMock {
	return &HistoryRepositoryMock{}
}

func (rep *HistoryRepositoryMock) GetAll() ([]models.History, error) {
	args := rep.Called()

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).([]models.History)

	if isSuccess {
		return data, nil
	} else {
		return []models.History{}, errors.New("testing error")
	}
}
func (rep *HistoryRepositoryMock) GetById(id uint) (models.History, error) {
	args := rep.Called(id)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(models.History)

	if isSuccess {
		return data, nil
	} else {
		return models.History{}, errors.New("testing error")
	}
}
func (rep *HistoryRepositoryMock) GetByPatientId(patientId uint) ([]models.History, error) {
	args := rep.Called(patientId)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).([]models.History)

	if isSuccess {
		return data, nil
	} else {
		return []models.History{}, errors.New("testing error")
	}
}
func (rep *HistoryRepositoryMock) GetByDoctorId(doctorId uint) ([]models.History, error) {
	args := rep.Called(doctorId)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).([]models.History)

	if isSuccess {
		return data, nil
	} else {
		return []models.History{}, errors.New("testing error")
	}
}
func (rep *HistoryRepositoryMock) Create(payload models.History) (models.History, error) {
	args := rep.Called(payload)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(models.History)

	if isSuccess {
		return data, nil
	} else {
		return models.History{}, errors.New("testing error")
	}
}
