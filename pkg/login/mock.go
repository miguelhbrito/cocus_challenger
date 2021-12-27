package login

import (
	"github.com/cocus_challenger/pkg/api/entity"
	"github.com/cocus_challenger/pkg/api/response"
)

type LoginCustomMock struct {
	CreateUserMock func(l entity.LoginEntity) error
	LoginMock      func(l entity.LoginEntity) (response.Token, error)
}

func (lm LoginCustomMock) CreateUser(l entity.LoginEntity) error {
	return lm.CreateUserMock(l)
}

func (lm LoginCustomMock) Login(l entity.LoginEntity) (response.Token, error) {
	return lm.LoginMock(l)
}
