package services

import "Tienda/entities"

type Service interface {
	FindAll() ([]entities.Product, error)

	Create(pro *entities.Product) bool

	Delete(id int64) bool
}
