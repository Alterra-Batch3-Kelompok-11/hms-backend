package mock

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"hms-backend/models"
)

type DoctorRepositoryMock struct {
	mock.Mock
}

func New() *DoctorRepositoryMock {
	return &DoctorRepositoryMock{}
}

func (rep *DoctorRepositoryMock) GetAll() ([]models.Doctor, error) {
	args := rep.Called()

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).([]models.Doctor)

	if isSuccess {
		return data, nil
	} else {
		return []models.Doctor{}, errors.New("testing error")
	}
}
func (rep *DoctorRepositoryMock) GetById(id uint) (models.Doctor, error) {
	args := rep.Called(id)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(models.Doctor)

	if isSuccess {
		return data, nil
	} else {
		return models.Doctor{}, errors.New("testing error")
	}
}
func (rep *DoctorRepositoryMock) GetByUserId(userId uint) (models.Doctor, error) {
	args := rep.Called(userId)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(models.Doctor)

	if isSuccess {
		return data, nil
	} else {
		return models.Doctor{}, errors.New("testing error")
	}
}
func (rep *DoctorRepositoryMock) GetByLicenseNumber(licenseNumber string) (models.Doctor, error) {
	args := rep.Called(licenseNumber)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(models.Doctor)

	if isSuccess {
		return data, nil
	} else {
		return models.Doctor{}, errors.New("testing error")
	}
}
func (rep *DoctorRepositoryMock) GetByLicenseNumberOther(licenseNumber string, id uint) (models.Doctor, error) {
	args := rep.Called(licenseNumber, id)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(models.Doctor)

	if isSuccess {
		return data, nil
	} else {
		return models.Doctor{}, errors.New("testing error")
	}
}
func (rep *DoctorRepositoryMock) GetBySpecialityId(specialityId uint) ([]models.Doctor, error) {
	args := rep.Called(specialityId)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).([]models.Doctor)

	if isSuccess {
		return data, nil
	} else {
		return []models.Doctor{}, errors.New("testing error")
	}
}
func (rep *DoctorRepositoryMock) Count() (int64, error) {
	args := rep.Called()

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(int64)

	if isSuccess {
		return data, nil
	} else {
		return int64(0), errors.New("testing error")
	}
}
func (rep *DoctorRepositoryMock) Create(user models.Doctor) (models.Doctor, error) {
	args := rep.Called(user)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(models.Doctor)

	if isSuccess {
		return data, nil
	} else {
		return models.Doctor{}, errors.New("testing error")
	}
}
func (rep *DoctorRepositoryMock) Update(id uint, user models.Doctor) (models.Doctor, error) {
	args := rep.Called(id, user)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(models.Doctor)

	if isSuccess {
		return data, nil
	} else {
		return models.Doctor{}, errors.New("testing error")
	}
}
func (rep *DoctorRepositoryMock) Delete(id uint) error {
	args := rep.Called(id)

	isSuccess := args.Get(0).(bool)

	if isSuccess {
		return nil
	} else {
		return errors.New("testing error")
	}
}
