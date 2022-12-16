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

func (uc *DoctorRepositoryMock) GetAll() ([]models.Doctor, error) {
	args := uc.Called()

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).([]models.Doctor)

	if isSuccess {
		return data, nil
	} else {
		return []models.Doctor{}, errors.New("testing error")
	}
}
func (uc *DoctorRepositoryMock) GetById(id uint) (models.Doctor, error) {
	args := uc.Called(id)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(models.Doctor)

	if isSuccess {
		return data, nil
	} else {
		return models.Doctor{}, errors.New("testing error")
	}
}
func (uc *DoctorRepositoryMock) GetByUserId(userId uint) (models.Doctor, error) {
	args := uc.Called(userId)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(models.Doctor)

	if isSuccess {
		return data, nil
	} else {
		return models.Doctor{}, errors.New("testing error")
	}
}
func (uc *DoctorRepositoryMock) GetByLicenseNumber(licenseNumber string) (models.Doctor, error) {
	args := uc.Called(licenseNumber)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(models.Doctor)

	if isSuccess {
		return data, nil
	} else {
		return models.Doctor{}, errors.New("testing error")
	}
}
func (uc *DoctorRepositoryMock) GetByLicenseNumberOther(licenseNumber string, id uint) (models.Doctor, error) {
	args := uc.Called(licenseNumber, id)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(models.Doctor)

	if isSuccess {
		return data, nil
	} else {
		return models.Doctor{}, errors.New("testing error")
	}
}
func (uc *DoctorRepositoryMock) GetBySpecialityId(specialityId uint) ([]models.Doctor, error) {
	args := uc.Called(specialityId)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).([]models.Doctor)

	if isSuccess {
		return data, nil
	} else {
		return []models.Doctor{}, errors.New("testing error")
	}
}
func (uc *DoctorRepositoryMock) Count() (int64, error) {
	args := uc.Called()

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(int64)

	if isSuccess {
		return data, nil
	} else {
		return int64(0), errors.New("testing error")
	}
}
func (uc *DoctorRepositoryMock) Create(user models.Doctor) (models.Doctor, error) {
	args := uc.Called(user)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(models.Doctor)

	if isSuccess {
		return data, nil
	} else {
		return models.Doctor{}, errors.New("testing error")
	}
}
func (uc *DoctorRepositoryMock) Update(id uint, user models.Doctor) (models.Doctor, error) {
	args := uc.Called(id, user)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(models.Doctor)

	if isSuccess {
		return data, nil
	} else {
		return models.Doctor{}, errors.New("testing error")
	}
}
func (uc *DoctorRepositoryMock) Delete(id uint) error {
	args := uc.Called(id)

	isSuccess := args.Get(0).(bool)

	if isSuccess {
		return nil
	} else {
		return errors.New("testing error")
	}
}
