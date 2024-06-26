package entity

import (
	"errors"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}

func (p *Product) UpdateName(name string) error {
	if len(name) == 0 {
		// エラー分はキャピタライずしちゃだめ？
		return errors.New("Name is required")
	}
	p.Name = name
	return nil
}
