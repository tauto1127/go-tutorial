package database

import "go-clean-architecture-tutorial/core/entity"
import "go-clean-architecture-tutorial/core/repository"

type ProductDatabaseArray struct {
	repository.ProductRepository
	products []entity.Product
}

func (p *ProductDatabaseArray) FindByID(id int) (*entity.Product, error) {
	for _, product := range p.products {
		if product.ID == id {
			return &product, nil
		}
	}
	return nil, nil
}

func (p *ProductDatabaseArray) Save(product *entity.Product) error {
	p.products = append(p.products, *product)
	return nil
}

// #TODO コントローラーとユースケースを定義しなきゃダメかも
func (p *ProductDatabaseArray) GetAll() ([]entity.Product, error) {
	return p.products, nil
}
