package repository

import (
	"Tienda/entities"
)

type ProductRepositoryMock struct{}

func NewProductRepositoryMock() ProductRepositoryMock {
	return ProductRepositoryMock{}
}

func (pr ProductRepositoryMock) FindAll() ([]entities.Product, error) {
	products := []entities.Product{
		{Id: 20, Name: "Glade", Price: 10.3, Quantity: 10, Status: true},
		{Id: 20, Name: "Glade", Price: 10.3, Quantity: 10, Status: true},
	}
	return products, nil
}

func (pl ProductRepositoryMock) Create(pro *entities.Product) bool {

	return false
}

func (pr ProductRepositoryMock) Delete(id int64) bool {
	return false
}
