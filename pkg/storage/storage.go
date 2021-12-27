package storage

import "github.com/cocus_challenger/pkg/api/entity"

type Triangle interface {
	Save(t entity.Triangle) error
	List() ([]entity.Triangle, error)
}

type Login interface {
	Save(l entity.LoginEntity) error
	Login(l entity.LoginEntity) (entity.LoginEntity, error)
}
