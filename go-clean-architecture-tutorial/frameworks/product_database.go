package frameworks

import "go-clean-architecture-tutorial/core/entity"
import "go-clean-architecture-tutorial/core/repository"

var products = []entity.Product{
	entity.Product{ID: 1, Name: "Product 1", Price: 100},
	entity.Product{ID: 2, Name: "Product 2", Price: 200},
}

type ProductDatabaseArray struct {
	repository.ProductRepository
}

func (p *ProductDatabaseArray) FindByID(id int) (*entity.Product, error) {
	for _, product := range products {
		if product.ID == id {
			return &product, nil
		}
	}
	return nil, nil
}

func (p *ProductDatabaseArray) Save(product *entity.Product) error {
	products = append(products, *product)
	return nil
}

func (p *ProductDatabaseArray) GetAll() ([]entity.Product, error) {
	return products, nil
}
