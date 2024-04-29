package entity

import (
	"errors"
)

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func (p *Product) UpdateName(name string) error {
	if len(name) == 0 {
		// エラー分はキャピタライずしちゃだめ？
		return errors.New("Name is required")
	}
	p.Name = name
	return nil
}
