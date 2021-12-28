package login

import (
	"errors"

	"github.com/golang-jwt/jwt"
)

var (
	JwtKey                 = []byte("my_secret_key")
	ErrUserOrPassIncorrect = errors.New("Username or Password is incorrect")
	ErrPasswordHash        = errors.New("Error to generate password hash")
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type Login struct {
	Username string `db:"username" json:"id"`   // Unique username.
	Password string `db:"password" json:"name"` // Hashed password.
}

type NewLogin struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Token struct {
	Token   string `json:"acess_token"`
	ExpTime int64  `json:"expTime"`
}

func (l NewLogin) GenerateEntity() Login {
	return Login{
		Username: l.Username,
		Password: l.Password,
	}
}

func (l NewLogin) Validate() error {
	var errs = ""
	if l.Username == "" {
		errs += "username is required"
	}
	if l.Password == "" {
		errs += ",password is required"
	}
	if len(errs) > 0 {
		return errors.New(errs)
	}
	return nil
}
