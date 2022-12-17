package mock

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"hms-backend/models"
	"time"
)

type OutpatientSessionRepositoryMock struct {
	mock.Mock
}

func New() *OutpatientSessionRepositoryMock {
	return &OutpatientSessionRepositoryMock{}
}

func (rep *OutpatientSessionRepositoryMock) GetAll() ([]models.OutpatientSession, error) {
	args := rep.Called()

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).([]models.OutpatientSession)

	if isSuccess {
		return data, nil
	} else {
		return []models.OutpatientSession{}, errors.New("testing error")
	}
}
func (rep *OutpatientSessionRepositoryMock) GetDesc(limit int) ([]models.OutpatientSession, error) {
	args := rep.Called(limit)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).([]models.OutpatientSession)

	if isSuccess {
		return data, nil
	} else {
		return []models.OutpatientSession{}, errors.New("testing error")
	}
}
func (rep *OutpatientSessionRepositoryMock) GetById(id uint) (models.OutpatientSession, error) {
	args := rep.Called(id)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(models.OutpatientSession)

	if isSuccess {
		return data, nil
	} else {
		return models.OutpatientSession{}, errors.New("testing error")
	}
}
func (rep *OutpatientSessionRepositoryMock) GetByDoctorId(doctorId uint) ([]models.OutpatientSession, error) {
	args := rep.Called(doctorId)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).([]models.OutpatientSession)

	if isSuccess {
		return data, nil
	} else {
		return []models.OutpatientSession{}, errors.New("testing error")
	}
}
func (rep *OutpatientSessionRepositoryMock) GetByPatientId(patientId uint) ([]models.OutpatientSession, error) {
	args := rep.Called(patientId)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).([]models.OutpatientSession)

	if isSuccess {
		return data, nil
	} else {
		return []models.OutpatientSession{}, errors.New("testing error")
	}
}
func (rep *OutpatientSessionRepositoryMock) GetUnprocessedByDoctorId(doctorId uint) ([]models.OutpatientSession, error) {
	args := rep.Called(doctorId)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).([]models.OutpatientSession)

	if isSuccess {
		return data, nil
	} else {
		return []models.OutpatientSession{}, errors.New("testing error")
	}
}
func (rep *OutpatientSessionRepositoryMock) GetProcessedByDoctorId(doctorId uint) ([]models.OutpatientSession, error) {
	args := rep.Called(doctorId)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).([]models.OutpatientSession)

	if isSuccess {
		return data, nil
	} else {
		return []models.OutpatientSession{}, errors.New("testing error")
	}
}
func (rep *OutpatientSessionRepositoryMock) GetProcessedAllByDoctorId(doctorId uint) ([]models.OutpatientSession, error) {
	args := rep.Called(doctorId)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).([]models.OutpatientSession)

	if isSuccess {
		return data, nil
	} else {
		return []models.OutpatientSession{}, errors.New("testing error")
	}
}
func (rep *OutpatientSessionRepositoryMock) GetApprovedByDoctorId(doctorId uint) ([]models.OutpatientSession, error) {
	args := rep.Called(doctorId)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).([]models.OutpatientSession)

	if isSuccess {
		return data, nil
	} else {
		return []models.OutpatientSession{}, errors.New("testing error")
	}
}
func (rep *OutpatientSessionRepositoryMock) GetApprovedAllByDoctorId(doctorId uint) ([]models.OutpatientSession, error) {
	args := rep.Called(doctorId)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).([]models.OutpatientSession)

	if isSuccess {
		return data, nil
	} else {
		return []models.OutpatientSession{}, errors.New("testing error")
	}
}
func (rep *OutpatientSessionRepositoryMock) GetRejectedByDoctorId(doctorId uint) ([]models.OutpatientSession, error) {
	args := rep.Called(doctorId)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).([]models.OutpatientSession)

	if isSuccess {
		return data, nil
	} else {
		return []models.OutpatientSession{}, errors.New("testing error")
	}
}
func (rep *OutpatientSessionRepositoryMock) GetByDate(date time.Time) ([]models.OutpatientSession, error) {
	args := rep.Called(date)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).([]models.OutpatientSession)

	if isSuccess {
		return data, nil
	} else {
		return []models.OutpatientSession{}, errors.New("testing error")
	}
}
func (rep *OutpatientSessionRepositoryMock) GetUnfinishedByDateByDoctorId(doctorId uint, date time.Time) ([]models.OutpatientSession, error) {
	args := rep.Called(doctorId, date)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).([]models.OutpatientSession)

	if isSuccess {
		return data, nil
	} else {
		return []models.OutpatientSession{}, errors.New("testing error")
	}
}
func (rep *OutpatientSessionRepositoryMock) GetFinishedByPatientIdDesc(patientId uint) ([]models.OutpatientSession, error) {
	args := rep.Called(patientId)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).([]models.OutpatientSession)

	if isSuccess {
		return data, nil
	} else {
		return []models.OutpatientSession{}, errors.New("testing error")
	}
}
func (rep *OutpatientSessionRepositoryMock) CountUnfinishedToday(doctorId uint, date time.Time) (int64, error) {
	args := rep.Called(doctorId, date)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(int64)

	if isSuccess {
		return data, nil
	} else {
		return 0, errors.New("testing error")
	}
}
func (rep *OutpatientSessionRepositoryMock) CountFinishedToday(doctorId uint, date time.Time) (int64, error) {
	args := rep.Called(doctorId, date)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(int64)

	if isSuccess {
		return data, nil
	} else {
		return 0, errors.New("testing error")
	}
}
func (rep *OutpatientSessionRepositoryMock) Create(user models.OutpatientSession) (models.OutpatientSession, error) {
	args := rep.Called(user)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(models.OutpatientSession)

	if isSuccess {
		return data, nil
	} else {
		return models.OutpatientSession{}, errors.New("testing error")
	}
}
func (rep *OutpatientSessionRepositoryMock) Update(id uint, user models.OutpatientSession) (models.OutpatientSession, error) {
	args := rep.Called(id, user)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(models.OutpatientSession)

	if isSuccess {
		return data, nil
	} else {
		return models.OutpatientSession{}, errors.New("testing error")
	}
}
func (rep *OutpatientSessionRepositoryMock) Approval(id uint, isApproved int) (models.OutpatientSession, error) {
	args := rep.Called(id, isApproved)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(models.OutpatientSession)

	if isSuccess {
		return data, nil
	} else {
		return models.OutpatientSession{}, errors.New("testing error")
	}
}
func (rep *OutpatientSessionRepositoryMock) Delete(id uint) error {
	args := rep.Called(id)

	isSuccess := args.Get(0).(bool)

	if isSuccess {
		return nil
	} else {
		return errors.New("testing error")
	}
}
