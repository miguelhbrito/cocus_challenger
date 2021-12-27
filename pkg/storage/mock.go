package storage

import "github.com/cocus_challenger/pkg/api/entity"

type TriangleCustomMock struct {
	SaveMock func(t entity.Triangle) error
	ListMock func() ([]entity.Triangle, error)
}

type LoginCustomMock struct {
	SaveMock  func(l entity.LoginEntity) error
	LoginMock func(l entity.LoginEntity) (entity.LoginEntity, error)
}

func (tm TriangleCustomMock) Save(t entity.Triangle) error {
	return tm.SaveMock(t)
}

func (tm TriangleCustomMock) List() ([]entity.Triangle, error) {
	return tm.ListMock()
}

func (lm LoginCustomMock) Save(l entity.LoginEntity) error {
	return lm.SaveMock(l)
}

func (lm LoginCustomMock) Login(l entity.LoginEntity) (entity.LoginEntity, error) {
	return lm.LoginMock(l)
}
