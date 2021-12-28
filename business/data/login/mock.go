package login

type LoginIntCustomMock struct {
	SaveMock  func(l Login) error
	LoginMock func(l Login) (Login, error)
}

func (lm LoginIntCustomMock) Save(l Login) error {
	return lm.SaveMock(l)
}

func (lm LoginIntCustomMock) Login(l Login) (Login, error) {
	return lm.LoginMock(l)
}
