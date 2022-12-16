package mock

import "github.com/stretchr/testify/mock"

type UserRepositoryMock struct {
	mock.Mock
}

func New() *UserRepositoryMock {
	return &UserRepositoryMock{}
}
