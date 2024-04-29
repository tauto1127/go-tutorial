package repository

import "go-clean-architecture-tutorial/core/entity"

type CartRepository interface {
	FindByID(id int) (*entity.Cart, error)
	Save(c *entity.Cart) error
}
