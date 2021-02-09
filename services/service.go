package services

import (
	"Tienda/entities"
	r "Tienda/repository"

	"github.com/go-kit/kit/log"
)

type service struct {
	repo   r.Repository
	logger log.Logger
}

func NewService(pr r.Repository, pl log.Logger) service {
	return service{
		repo:   pr,
		logger: pl,
	}
}


func (s service) FindAll() ([]entities.Product, error) { //slice =lista
	//logger := log.With(s.logger, "metodo", "FindAll")
	products, _ := s.repo.FindAll()
	return products, nil
}

func (s service) Create(p *entities.Product) bool { //slice =lista
	//logger := log.With(s.logger, "metodo", "FindAll")
	band := s.repo.Create(p)
	return band
}

func (s service) Delete(id int64) bool { //slice =lista
	//logger := log.With(s.logger, "metodo", "FindAll")
	products := s.Delete(id)
	return products
}
