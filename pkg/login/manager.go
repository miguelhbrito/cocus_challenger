package login

import (
	"time"

	"github.com/cocus_challenger/pkg/api/entity"
	"github.com/cocus_challenger/pkg/api/response"
	"github.com/cocus_challenger/pkg/auth"
	"github.com/cocus_challenger/pkg/storage"
	"github.com/golang-jwt/jwt"
)

type manager struct {
	loginStorage storage.Login
	auth         auth.Auth
}

func NewManager(loginStorage storage.Login, auth auth.Auth) Login {
	return manager{
		loginStorage: loginStorage,
		auth:         auth,
	}
}

func (m manager) CreateUser(l entity.LoginEntity) error {

	//Generation new hashedpassword to save into db
	newPassword, err := m.auth.GenerateHashPassword(l.Password)
	if err != nil {
		return errPasswordHash
	}

	//Saving user into db
	l.Password = newPassword
	err = m.loginStorage.Save(l)
	if err != nil {
		return err
	}
	return nil
}

func (m manager) Login(l entity.LoginEntity) (response.Token, error) {
	//Getting credentials from database
	lr, err := m.loginStorage.Login(l)
	if err != nil {
		return response.Token{}, err
	}

	//Checking input secretHash with secretHash from database
	check := m.auth.CheckPasswordHash(l.Password, lr.Password)
	if !check {
		return response.Token{}, errUserOrPassIncorrect
	}

	//Generation jwt token
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &Claims{
		Username: l.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//Signing jwt token with our key
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		return response.Token{}, err
	}

	tokenResponse := response.Token{
		Token:   tokenString,
		ExpTime: expirationTime.Unix(),
	}

	return tokenResponse, err
}
