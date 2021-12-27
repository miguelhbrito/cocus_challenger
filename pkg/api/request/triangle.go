package request

import (
	"errors"

	"github.com/cocus_challenger/pkg/api/entity"
	"github.com/google/uuid"
)

type Triangle struct {
	Side1 int `json:"side1"`
	Side2 int `json:"side2"`
	Side3 int `json:"side3"`
}

func (t Triangle) GenerateEntity() entity.Triangle {
	return entity.Triangle{
		Id:    uuid.New().String(),
		Side1: t.Side1,
		Side2: t.Side2,
		Side3: t.Side3,
	}
}

func (t Triangle) Validate() error {
	var errs = ""
	if t.Side1 <= 0 {
		errs += "side1 can't be lower than 0 or equal 0"
	}
	if t.Side2 <= 0 {
		errs += ",side2 can't be lower than 0 or equal 0"
	}
	if t.Side3 <= 0 {
		errs += ",side3 can't be lower than 0 or equal 0"
	}
	if len(errs) > 0 {
		return errors.New(errs)
	}
	return nil
}
