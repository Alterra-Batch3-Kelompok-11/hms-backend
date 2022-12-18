package mock

import (
	"errors"
	"hms-backend/dto"

	"github.com/stretchr/testify/mock"
)

type DoctorScheduleUseCaseMock struct {
	mock.Mock
}

func New() *DoctorScheduleUseCaseMock {
	return &DoctorScheduleUseCaseMock{}
}

func (uc *DoctorScheduleUseCaseMock) GetById(id uint) (dto.DoctorScheduleRes, error) {
	args := uc.Called(id)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(dto.DoctorScheduleRes)

	if isSuccess {
		return data, nil
	} else {
		return dto.DoctorScheduleRes{}, errors.New("testing error")
	}
}

func (uc *DoctorScheduleUseCaseMock) GetByDoctorId(doctorId uint) ([]dto.DoctorScheduleRes, error) {
	args := uc.Called(doctorId)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).([]dto.DoctorScheduleRes)

	if isSuccess {
		return data, nil
	} else {
		return []dto.DoctorScheduleRes{}, errors.New("testing error")
	}
}

func (uc *DoctorScheduleUseCaseMock) GetByLicenseNumber(licenseNumber string) ([]dto.DoctorScheduleRes, error) {
	args := uc.Called(licenseNumber)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).([]dto.DoctorScheduleRes)

	if isSuccess {
		return data, nil
	} else {
		return []dto.DoctorScheduleRes{}, errors.New("testing error")
	}
}

func (uc *DoctorScheduleUseCaseMock) Create(payload dto.DoctorScheduleReq) (dto.DoctorScheduleRes, error) {
	args := uc.Called(payload)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(dto.DoctorScheduleRes)

	if isSuccess {
		return data, nil
	} else {
		return dto.DoctorScheduleRes{}, errors.New("testing error")
	}
}

func (uc *DoctorScheduleUseCaseMock) Update(id uint, payload dto.DoctorScheduleReq) (dto.DoctorScheduleRes, error) {
	args := uc.Called(id, payload)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(dto.DoctorScheduleRes)

	if isSuccess {
		return data, nil
	} else {
		return dto.DoctorScheduleRes{}, errors.New("testing error")
	}
}

func (uc *DoctorScheduleUseCaseMock) Delete(id uint) error {
	args := uc.Called(id)

	isSuccess := args.Get(0).(bool)

	if isSuccess {
		return nil
	} else {
		return errors.New("testing error")
	}
}
