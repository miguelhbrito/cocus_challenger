package triangle

import "github.com/cocus_challenger/pkg/api/entity"

type TriangleCustomMock struct {
	CreateMock func(t entity.Triangle) (entity.Triangle, error)
	ListMock   func() (entity.Triangles, error)
}

func (tm TriangleCustomMock) Create(t entity.Triangle) (entity.Triangle, error) {
	return tm.CreateMock(t)
}

func (tm TriangleCustomMock) List() (entity.Triangles, error) {
	return tm.ListMock()
}
