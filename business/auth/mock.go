package auth

type AuthCustomMock struct {
	CheckPasswordHashMock    func(password, hash string) bool
	GenerateHashPasswordMock func(password string) (string, error)
}

func (a AuthCustomMock) CheckPasswordHash(password, hash string) bool {
	return a.CheckPasswordHashMock(password, hash)
}

func (a AuthCustomMock) GenerateHashPassword(password string) (string, error) {
	return a.GenerateHashPasswordMock(password)
}
