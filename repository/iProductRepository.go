package repository

import "Tienda/entities"

type Repository interface {
	FindAll() ([]entities.Product, error)

	Create(pro *entities.Product) bool

	Delete(id int64) bool
}
