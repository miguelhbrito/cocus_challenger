package triangle

import (
	"errors"

	"github.com/google/uuid"
)

const (
	Equilateral string = "equilateral"
	Isosceles          = "isosceles"
	Scalene            = "scalene"
)

var (
	ErrNotATriangle = errors.New("is not a triangle, the sum of two sides is greater than the other side")
)

type Triangle struct {
	Id    string `db:"id" json:"id"`
	Side1 int    `db:"side1" json:"side1"`
	Side2 int    `db:"side2" json:"side2"`
	Side3 int    `db:"side3" json:"side3"`
	Type  string `db:"type" json:"type"`
}

type NewTriangle struct {
	Side1 int `json:"side1"`
	Side2 int `json:"side2"`
	Side3 int `json:"side3"`
}

type Triangles []Triangle

func (t NewTriangle) GenerateEntity() Triangle {
	return Triangle{
		Id:    uuid.New().String(),
		Side1: t.Side1,
		Side2: t.Side2,
		Side3: t.Side3,
	}
}

func (t Triangle) Response() Triangle {
	return Triangle{
		Id:    t.Id,
		Side1: t.Side1,
		Side2: t.Side2,
		Side3: t.Side3,
		Type:  t.Type,
	}
}

func (t Triangles) Response() []Triangle {
	resp := make([]Triangle, 0)
	for i := range t {
		resp = append(resp, t[i].Response())
	}
	return resp
}

func (t NewTriangle) Validate() error {
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
