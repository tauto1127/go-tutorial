package usecase

import (
	"go-clean-architecture-tutorial/core/repository"
)

type AddProductToCartUsecase struct {
	ProductRepository repository.ProductRepository
	CartRepository    repository.CartRepository
}

// Execute adds a Product to a Cart
func (uc *AddProductToCartUsecase) Execute(productID int, cartID int) error {
	// Get the Product from the ProductRepository
	product, err := uc.ProductRepository.FindByID(productID)
	if err != nil {
		return err
	}

	// Get the Cart from the CartRepository
	cart, err := uc.CartRepository.FindByID(cartID)
	if err != nil {
		return err
	}

	// Add the Product to the Cart
	err = cart.AddProduct(product)
	if err != nil {
		return err
	}

	// Save the updated Cart to the CartRepository
	err = uc.CartRepository.Save(cart)
	if err != nil {
		return err
	}

	return nil
}
