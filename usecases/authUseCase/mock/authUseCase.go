package mock

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"hms-backend/dto"
)

type AuthUseCaseMock struct {
	mock.Mock
}

func New() *AuthUseCaseMock {
	return &AuthUseCaseMock{}
}

func (uc *AuthUseCaseMock) Login(username, password string) (dto.LoginRes, error) {
	args := uc.Called(username, password)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(dto.LoginRes)

	if isSuccess {
		return data, nil
	} else {
		return dto.LoginRes{}, errors.New("testing error")
	}
}
func (uc *AuthUseCaseMock) SignUp(payload dto.UserReq) (dto.UserRes, error) {
	args := uc.Called(payload)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(dto.UserRes)

	if isSuccess {
		return data, nil
	} else {
		return dto.UserRes{}, errors.New("testing error")
	}
}
func (uc *AuthUseCaseMock) RefreshToken(id uint) (dto.LoginRes, error) {
	args := uc.Called(id)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(dto.LoginRes)

	if isSuccess {
		return data, nil
	} else {
		return dto.LoginRes{}, errors.New("testing error")
	}
}
