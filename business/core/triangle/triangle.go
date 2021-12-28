package triangle

import (
	"github.com/cocus_challenger_refact/business/data/triangle"
)

type TriangleInt interface {
	Create(t triangle.Triangle) (triangle.Triangle, error)
	List() (triangle.Triangles, error)
}

type Core struct {
	db triangle.TriangleInt
}

func NewCore(db triangle.TriangleInt) TriangleInt {
	return Core{
		db: db,
	}
}

func (c Core) Create(t triangle.Triangle) (triangle.Triangle, error) {

	//Check if is a valid triangle
	if !isTriangle(t) {
		return triangle.Triangle{}, triangle.ErrNotATriangle
	}

	//Check which type is the triangle
	if allSidesAreEqual(t) {
		t.Type = triangle.Equilateral
	} else if twoSidesAreEqual(t) {
		t.Type = triangle.Isosceles
	} else {
		t.Type = triangle.Scalene
	}

	//Save triangle into db
	err := c.db.Save(t)
	if err != nil {
		return triangle.Triangle{}, err
	}

	return t, nil
}

func (c Core) List() (triangle.Triangles, error) {

	triangles, err := c.db.List()
	if err != nil {
		return nil, err
	}

	return triangles, nil
}

func isTriangle(t triangle.Triangle) bool {
	if t.Side1+t.Side2 <= t.Side3 || t.Side1+t.Side3 <= t.Side2 || t.Side2+t.Side3 <= t.Side1 {
		return false
	} else {
		return true
	}
}

func allSidesAreEqual(t triangle.Triangle) bool {
	if t.Side1 == t.Side2 && t.Side2 == t.Side3 {
		return true
	}
	return false
}

func twoSidesAreEqual(t triangle.Triangle) bool {
	if t.Side1 == t.Side2 || t.Side1 == t.Side3 || t.Side2 == t.Side3 {
		return true
	}
	return false
}
