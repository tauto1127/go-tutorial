package controllers

import (
	"go-clean-architecture-tutorial/core/entity"
	"go-clean-architecture-tutorial/core/usecase"
	"net/http"
	"strconv"
)

type AddProductToCartController struct {
	UseCase usecase.AddProductToCartUsecase
}

func (c *AddProductToCartController) Handle(r *http.Request) error {
	productID, err := strconv.Atoi(r.FormValue("product_id"))
	if err != nil {
		return err
	}

	cartID, err := strconv.Atoi(r.FormValue("cart_id"))
	if err != nil {
		return err
	}

	err = c.UseCase.Execute(productID, cartID)
	if err != nil {
		return err
	}

	return nil
}

type GetAllProductsController struct {
	UseCase usecase.GetAllProductsUsecase
}

func (c *GetAllProductsController) Handle(r *http.Request) (products []entity.Product, error error) {
	products, err := c.UseCase.Execute()
	if err != nil {
		return
	}
	return products, nil
}
