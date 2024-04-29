package repository

import "go-clean-architecture-tutorial/core/entity"

type ProductRepository interface {
	FindByID(id int) (*entity.Product, error)
	Save(p *entity.Product) error
	GetAll() ([]entity.Product, error)
}
