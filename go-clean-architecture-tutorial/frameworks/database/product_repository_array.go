package database

import "go-clean-architecture-tutorial/core/entity"

var ProductArray ProductDatabaseArray = ProductDatabaseArray{
	products: []entity.Product{
		{ID: 1, Name: "Product 1", Price: 100},
		{ID: 2, Name: "Product 2", Price: 200},
		{ID: 3, Name: "Product 3", Price: 300},
	},
}
