package presenter

import "go-clean-architecture-tutorial/core/entity"

type ProductPresenter struct{}

type ProductDTO struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func (p *ProductPresenter) Present(products []*entity.Product) []ProductDTO {
	var result []ProductDTO

	for _, product := range products {
		result = append(result, ProductDTO{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Quantity:    product.Quantity,
		})
	}

	return result
}
