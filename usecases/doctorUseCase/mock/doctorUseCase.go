package mock

import (
	"errors"
	"hms-backend/dto"

	"github.com/stretchr/testify/mock"
)

type DoctorUseCaseMock struct {
	mock.Mock
}

func New() *DoctorUseCaseMock {
	return &DoctorUseCaseMock{}
}

func (uc *DoctorUseCaseMock) GetAll() ([]dto.DoctorRes, error) {
	args := uc.Called()

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).([]dto.DoctorRes)

	if isSuccess {
		return data, nil
	} else {
		return []dto.DoctorRes{}, errors.New("testing error")
	}
}
func (uc *DoctorUseCaseMock) GetById(id uint) (dto.DoctorRes, error) {
	args := uc.Called(id)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(dto.DoctorRes)

	if isSuccess {
		return data, nil
	} else {
		return dto.DoctorRes{}, errors.New("testing error")
	}
}
func (uc *DoctorUseCaseMock) GetByLicenseNumber(licenseNumber string) (dto.DoctorRes, error) {
	args := uc.Called(licenseNumber)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(dto.DoctorRes)

	if isSuccess {
		return data, nil
	} else {
		return dto.DoctorRes{}, errors.New("testing error")
	}
}
func (uc *DoctorUseCaseMock) GetBySpecialityId(specialityId uint) ([]dto.DoctorRes, error) {
	args := uc.Called(specialityId)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).([]dto.DoctorRes)

	if isSuccess {
		return data, nil
	} else {
		return []dto.DoctorRes{}, errors.New("testing error")
	}
}
func (uc *DoctorUseCaseMock) GetToday() ([]dto.DoctorRes, error) {
	args := uc.Called()

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).([]dto.DoctorRes)

	if isSuccess {
		return data, nil
	} else {
		return []dto.DoctorRes{}, errors.New("testing error")
	}
}
func (uc *DoctorUseCaseMock) Create(payload dto.DoctorReq) (dto.DoctorRes, error) {
	args := uc.Called(payload)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(dto.DoctorRes)

	if isSuccess {
		return data, nil
	} else {
		return dto.DoctorRes{}, errors.New("testing error")
	}
}
func (uc *DoctorUseCaseMock) Update(id uint, payload dto.DoctorReq) (dto.DoctorRes, error) {
	args := uc.Called(id, payload)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(dto.DoctorRes)

	if isSuccess {
		return data, nil
	} else {
		return dto.DoctorRes{}, errors.New("testing error")
	}
}
func (uc *DoctorUseCaseMock) Delete(id uint) error {
	args := uc.Called(id)

	isSuccess := args.Get(0).(bool)

	if isSuccess {
		return nil
	} else {
		return errors.New("testing error")
	}
}
