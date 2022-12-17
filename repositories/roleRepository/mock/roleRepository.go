package mock

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"hms-backend/models"
)

type RoleRepositoryMock struct {
	mock.Mock
}

func New() *RoleRepositoryMock {
	return &RoleRepositoryMock{}
}

func (rep *RoleRepositoryMock) GetAll() ([]models.Role, error) {
	args := rep.Called()

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).([]models.Role)

	if isSuccess {
		return data, nil
	} else {
		return []models.Role{}, errors.New("testing error")
	}
}
func (rep *RoleRepositoryMock) GetById(id uint) (models.Role, error) {
	args := rep.Called(id)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(models.Role)

	if isSuccess {
		return data, nil
	} else {
		return models.Role{}, errors.New("testing error")
	}
}
