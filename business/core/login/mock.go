package login

import "github.com/cocus_challenger_refact/business/data/login"

type LoginCustomMock struct {
	CreateUserMock func(l login.Login) error
	LoginMock      func(l login.Login) (login.Token, error)
}

func (lm LoginCustomMock) CreateUser(l login.Login) error {
	return lm.CreateUserMock(l)
}

func (lm LoginCustomMock) Login(l login.Login) (login.Token, error) {
	return lm.LoginMock(l)
}
