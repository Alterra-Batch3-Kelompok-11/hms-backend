package mock

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"hms-backend/models"
)

type DoctorScheduleRepositoryMock struct {
	mock.Mock
}

func New() *DoctorScheduleRepositoryMock {
	return &DoctorScheduleRepositoryMock{}
}

func (rep *DoctorScheduleRepositoryMock) GetById(id uint) (models.DoctorSchedule, error) {
	args := rep.Called(id)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(models.DoctorSchedule)

	if isSuccess {
		return data, nil
	} else {
		return models.DoctorSchedule{}, errors.New("testing error")
	}
}
func (rep *DoctorScheduleRepositoryMock) GetByDoctorId(doctorId uint) ([]models.DoctorSchedule, error) {
	args := rep.Called(doctorId)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).([]models.DoctorSchedule)

	if isSuccess {
		return data, nil
	} else {
		return []models.DoctorSchedule{}, errors.New("testing error")
	}
}
func (rep *DoctorScheduleRepositoryMock) GetByDoctorIdDay(doctorId uint, dayInt int) (models.DoctorSchedule, error) {
	args := rep.Called(doctorId, dayInt)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(models.DoctorSchedule)

	if isSuccess {
		return data, nil
	} else {
		return models.DoctorSchedule{}, errors.New("testing error")
	}
}
func (rep *DoctorScheduleRepositoryMock) GetByDay(dayInt int) ([]models.DoctorSchedule, error) {
	args := rep.Called(dayInt)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).([]models.DoctorSchedule)

	if isSuccess {
		return data, nil
	} else {
		return []models.DoctorSchedule{}, errors.New("testing error")
	}
}
func (rep *DoctorScheduleRepositoryMock) Create(sched models.DoctorSchedule) (models.DoctorSchedule, error) {
	args := rep.Called(sched)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(models.DoctorSchedule)

	if isSuccess {
		return data, nil
	} else {
		return models.DoctorSchedule{}, errors.New("testing error")
	}
}
func (rep *DoctorScheduleRepositoryMock) Update(id uint, sched models.DoctorSchedule) (models.DoctorSchedule, error) {
	args := rep.Called(id, sched)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(models.DoctorSchedule)

	if isSuccess {
		return data, nil
	} else {
		return models.DoctorSchedule{}, errors.New("testing error")
	}
}
func (rep *DoctorScheduleRepositoryMock) Delete(id uint) error {
	args := rep.Called(id)

	isSuccess := args.Get(0).(bool)

	if isSuccess {
		return nil
	} else {
		return errors.New("testing error")
	}
}
