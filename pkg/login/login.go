package login

import (
	"errors"

	"github.com/cocus_challenger/pkg/api/entity"
	"github.com/cocus_challenger/pkg/api/response"
	"github.com/golang-jwt/jwt"
)

var (
	JwtKey                 = []byte("my_secret_key")
	errUserOrPassIncorrect = errors.New("Username or Password is incorrect")
	errPasswordHash        = errors.New("Error to generate password hash")
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type Login interface {
	CreateUser(l entity.LoginEntity) error
	Login(l entity.LoginEntity) (response.Token, error)
}
