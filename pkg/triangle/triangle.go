package triangle

import (
	"errors"

	"github.com/cocus_challenger/pkg/api/entity"
)

/*
* Scalene: No sides of the triangle are equal
* Isosceles: Any two sides of the triangle are equal
* Equilateral: All sides of the triangle are equal
 */

const (
	Equilateral string = "equilateral"
	Isosceles          = "isosceles"
	Scalene            = "scalene"
)

var (
	errNotATriangle = errors.New("is not a triangle, the sum of two sides is greater than the other side")
)

type Triangle interface {
	Create(t entity.Triangle) (entity.Triangle, error)
	List() (entity.Triangles, error)
}
