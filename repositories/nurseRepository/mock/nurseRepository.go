package mock

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"hms-backend/models"
)

type NurseRepositoryMock struct {
	mock.Mock
}

func New() *NurseRepositoryMock {
	return &NurseRepositoryMock{}
}

func (rep *NurseRepositoryMock) GetAll() ([]models.Nurse, error) {
	args := rep.Called()

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).([]models.Nurse)

	if isSuccess {
		return data, nil
	} else {
		return []models.Nurse{}, errors.New("testing error")
	}
}
func (rep *NurseRepositoryMock) GetById(id uint) (models.Nurse, error) {
	args := rep.Called(id)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(models.Nurse)

	if isSuccess {
		return data, nil
	} else {
		return models.Nurse{}, errors.New("testing error")
	}
}
func (rep *NurseRepositoryMock) GetByUserId(userId uint) (models.Nurse, error) {
	args := rep.Called(userId)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(models.Nurse)

	if isSuccess {
		return data, nil
	} else {
		return models.Nurse{}, errors.New("testing error")
	}
}
func (rep *NurseRepositoryMock) GetByLicenseNumber(licenseNumber string) (models.Nurse, error) {
	args := rep.Called(licenseNumber)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(models.Nurse)

	if isSuccess {
		return data, nil
	} else {
		return models.Nurse{}, errors.New("testing error")
	}
}
func (rep *NurseRepositoryMock) GetByLicenseNumberOther(licenseNumber string, id uint) (models.Nurse, error) {
	args := rep.Called(licenseNumber, id)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(models.Nurse)

	if isSuccess {
		return data, nil
	} else {
		return models.Nurse{}, errors.New("testing error")
	}
}
func (rep *NurseRepositoryMock) Count() (int64, error) {
	args := rep.Called()

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(int64)

	if isSuccess {
		return data, nil
	} else {
		return 0, errors.New("testing error")
	}
}
func (rep *NurseRepositoryMock) Create(user models.Nurse) (models.Nurse, error) {
	args := rep.Called(user)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(models.Nurse)

	if isSuccess {
		return data, nil
	} else {
		return models.Nurse{}, errors.New("testing error")
	}
}
func (rep *NurseRepositoryMock) Update(id uint, user models.Nurse) (models.Nurse, error) {
	args := rep.Called(id, user)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(models.Nurse)

	if isSuccess {
		return data, nil
	} else {
		return models.Nurse{}, errors.New("testing error")
	}
}
func (rep *NurseRepositoryMock) Delete(id uint) error {
	args := rep.Called(id)

	isSuccess := args.Get(0).(bool)

	if isSuccess {
		return nil
	} else {
		return errors.New("testing error")
	}
}
