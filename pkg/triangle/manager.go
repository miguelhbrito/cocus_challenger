package triangle

import (
	"github.com/cocus_challenger/pkg/api/entity"
	"github.com/cocus_challenger/pkg/storage"
)

type manager struct {
	triangleStorage storage.Triangle
}

func NewManager(triangleStorage storage.Triangle) Triangle {
	return manager{
		triangleStorage: triangleStorage,
	}
}

func (m manager) Create(t entity.Triangle) (entity.Triangle, error) {

	//Check if is a valid triangle
	if !isTriangle(t) {
		return entity.Triangle{}, errNotATriangle
	}

	//Check which type is the triangle
	if AllSidesAreEqual(t) {
		t.Type = Equilateral
	} else if TwoSidesAreEqual(t) {
		t.Type = Isosceles
	} else {
		t.Type = Scalene
	}

	//Save triangle into db
	err := m.triangleStorage.Save(t)
	if err != nil {
		return entity.Triangle{}, err
	}

	return t, nil
}

func (m manager) List() (entity.Triangles, error) {

	triangles, err := m.triangleStorage.List()
	if err != nil {
		return nil, err
	}

	return triangles, nil
}

func isTriangle(t entity.Triangle) bool {
	if t.Side1+t.Side2 <= t.Side3 || t.Side1+t.Side3 <= t.Side2 || t.Side2+t.Side3 <= t.Side1 {
		return false
	} else {
		return true
	}
}

func AllSidesAreEqual(t entity.Triangle) bool {
	if t.Side1 == t.Side2 && t.Side2 == t.Side3 {
		return true
	}
	return false
}

func TwoSidesAreEqual(t entity.Triangle) bool {
	if t.Side1 == t.Side2 || t.Side1 == t.Side3 || t.Side2 == t.Side3 {
		return true
	}
	return false
}
