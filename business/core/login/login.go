package login

import (
	"time"

	"github.com/cocus_challenger_refact/business/auth"
	"github.com/cocus_challenger_refact/business/data/login"
	"github.com/golang-jwt/jwt"
)

type LoginInt interface {
	CreateUser(l login.Login) error
	Login(l login.Login) (login.Token, error)
}

type Core struct {
	db   login.LoginInt
	auth auth.Auth
}

func NewCore(db login.LoginInt, auth auth.Auth) LoginInt {
	return Core{
		db:   db,
		auth: auth,
	}
}

func (c Core) CreateUser(l login.Login) error {

	//Generation new hashedpassword to save into db
	newPassword, err := c.auth.GenerateHashPassword(l.Password)
	if err != nil {
		return login.ErrPasswordHash
	}

	//Saving user into db
	l.Password = newPassword
	err = c.db.Save(l)
	if err != nil {
		return err
	}
	return nil
}

func (c Core) Login(l login.Login) (login.Token, error) {
	//Getting credentials from database
	lr, err := c.db.Login(l)
	if err != nil {
		return login.Token{}, err
	}

	//Checking input secretHash with secretHash from database
	check := c.auth.CheckPasswordHash(l.Password, lr.Password)
	if !check {
		return login.Token{}, login.ErrUserOrPassIncorrect
	}

	//Generation jwt token
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &login.Claims{
		Username: l.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//Signing jwt token with our key
	tokenString, err := token.SignedString(login.JwtKey)
	if err != nil {
		return login.Token{}, err
	}

	tokenResponse := login.Token{
		Token:   tokenString,
		ExpTime: expirationTime.Unix(),
	}

	return tokenResponse, err
}
