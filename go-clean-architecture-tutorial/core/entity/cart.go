package entity

import "errors"

type Cart struct {
	ID    int
	Items []CartItem
}

type CartItem struct {
	Product  *Product
	Quantity int
}

func (c *Cart) AddProduct(p *Product) error {
	if p.Quantity == 0 {
		return errors.New("Product quantity is zero")
	}

	// Check if the Product is already in the Cart
	for _, item := range c.Items {
		if item.Product.ID == p.ID {
			return errors.New("Product already exists in cart")
		}
	}
	// Add the Product to the Cart
	c.Items = append(c.Items, CartItem{
		Product:  p,
		Quantity: 1,
	})

	return nil
}
