package model

import uuid "github.com/satori/go.uuid"

type Product struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Products struct {
	Product []*Product `json:"products"`
}

func (p *Products) Add(product *Product) {
	p.Product = append(p.Product, product)
}

func NewProduct() *Product {
	return &Product{
		ID: uuid.NewV4().String(),
	}
}
