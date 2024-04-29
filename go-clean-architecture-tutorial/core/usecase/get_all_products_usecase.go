package usecase

import (
	"go-clean-architecture-tutorial/core/entity"
	"go-clean-architecture-tutorial/core/repository"
)

type GetAllProductsUsecase struct {
	ProductRepository repository.ProductRepository
}

// Execute returns all Products
func (uc *GetAllProductsUsecase) Execute() ([]entity.Product, error) {
	products, err := uc.ProductRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return products, nil
}
