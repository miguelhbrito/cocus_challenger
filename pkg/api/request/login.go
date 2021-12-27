package request

import (
	"errors"

	"github.com/cocus_challenger/pkg/api/entity"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (l LoginRequest) GenerateEntity() entity.LoginEntity {
	return entity.LoginEntity{
		Username: l.Username,
		Password: l.Password,
	}
}

func (l LoginRequest) Validate() error {
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
