package triangle

import "github.com/cocus_challenger_refact/business/data/triangle"

type TriangleCustomMock struct {
	CreateMock func(t triangle.Triangle) (triangle.Triangle, error)
	ListMock   func() (triangle.Triangles, error)
}

func (tm TriangleCustomMock) Create(t triangle.Triangle) (triangle.Triangle, error) {
	return tm.CreateMock(t)
}

func (tm TriangleCustomMock) List() (triangle.Triangles, error) {
	return tm.ListMock()
}
