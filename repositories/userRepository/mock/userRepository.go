package mock

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"hms-backend/models"
)

type UserRepositoryMock struct {
	mock.Mock
}

func New() *UserRepositoryMock {
	return &UserRepositoryMock{}
}

func (rep *UserRepositoryMock) GetAll() ([]models.User, error) {
	args := rep.Called()

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).([]models.User)

	if isSuccess {
		return data, nil
	} else {
		return []models.User{}, errors.New("testing error")
	}
}
func (rep *UserRepositoryMock) GetById(id uint) (models.User, error) {
	args := rep.Called(id)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(models.User)

	if isSuccess {
		return data, nil
	} else {
		return models.User{}, errors.New("testing error")
	}
}
func (rep *UserRepositoryMock) GetByUsernamePassword(username, password string) (models.User, error) {
	args := rep.Called(username, password)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(models.User)

	if isSuccess {
		return data, nil
	} else {
		return models.User{}, errors.New("testing error")
	}
}
func (rep *UserRepositoryMock) GetByUsername(username string) (models.User, error) {
	args := rep.Called(username)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(models.User)

	if isSuccess {
		return data, nil
	} else {
		return models.User{}, errors.New("testing error")
	}
}
func (rep *UserRepositoryMock) Create(user models.User) (models.User, error) {
	args := rep.Called(user)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(models.User)

	if isSuccess {
		return data, nil
	} else {
		return models.User{}, errors.New("testing error")
	}
}
func (rep *UserRepositoryMock) Update(id uint, user models.User) (models.User, error) {
	args := rep.Called(id, user)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(models.User)

	if isSuccess {
		return data, nil
	} else {
		return models.User{}, errors.New("testing error")
	}
}
func (rep *UserRepositoryMock) Delete(id uint) error {
	args := rep.Called(id)

	isSuccess := args.Get(0).(bool)

	if isSuccess {
		return nil
	} else {
		return errors.New("testing error")
	}
}
