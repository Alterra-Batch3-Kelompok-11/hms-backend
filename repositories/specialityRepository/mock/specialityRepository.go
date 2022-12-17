package mock

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"hms-backend/models"
)

type SpecialityRepositoryMock struct {
	mock.Mock
}

func New() *SpecialityRepositoryMock {
	return &SpecialityRepositoryMock{}
}

func (rep *SpecialityRepositoryMock) GetAll() ([]models.Speciality, error) {
	args := rep.Called()

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).([]models.Speciality)

	if isSuccess {
		return data, nil
	} else {
		return []models.Speciality{}, errors.New("testing error")
	}
}
func (rep *SpecialityRepositoryMock) GetById(id uint) (models.Speciality, error) {
	args := rep.Called(id)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(models.Speciality)

	if isSuccess {
		return data, nil
	} else {
		return models.Speciality{}, errors.New("testing error")
	}
}
func (rep *SpecialityRepositoryMock) Create(payload models.Speciality) (models.Speciality, error) {
	args := rep.Called(payload)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(models.Speciality)

	if isSuccess {
		return data, nil
	} else {
		return models.Speciality{}, errors.New("testing error")
	}
}
func (rep *SpecialityRepositoryMock) Update(id uint, payload models.Speciality) (models.Speciality, error) {
	args := rep.Called(id, payload)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(models.Speciality)

	if isSuccess {
		return data, nil
	} else {
		return models.Speciality{}, errors.New("testing error")
	}
}
func (rep *SpecialityRepositoryMock) Delete(id uint) error {
	args := rep.Called(id)

	isSuccess := args.Get(0).(bool)

	if isSuccess {
		return nil
	} else {
		return errors.New("testing error")
	}
}
