package config

import (
	e "Tienda/entities"
	s "Tienda/services"
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetAllProduct endpoint.Endpoint
	CreateProduct endpoint.Endpoint
	DeleteProduct endpoint.Endpoint
}

func CreateEndpoints(s s.Service) Endpoints {
	return Endpoints{
		GetAllProduct: MakeGetAllProductEnd(s),
		CreateProduct: MakeCreateProductEnd(s),
		DeleteProduct: MakeDeleteProductEnd(s),
	}
}

// Summary Return Response
// @Description findAll
// @Tags product
// @Accept  json
// @Produce  json
// @Success 200 {object} []entities.Product
// @Router /Tienda/findAll [get]
func MakeGetAllProductEnd(s s.Service) endpoint.Endpoint {
	return func(ct context.Context, r interface{}) (interface{}, error) {
		result, err := s.FindAll()
		return result, err
	}
}

func MakeCreateProductEnd(s s.Service) endpoint.Endpoint {
	return func(ct context.Context, r interface{}) (interface{}, error) {
		req := r.(e.Product)     //cast
		result := s.Create(&req) //la & por que recibe puntero   agregar error***************************
		return result, nil
	}
}

func MakeDeleteProductEnd(s s.Service) endpoint.Endpoint {
	return func(ct context.Context, r interface{}) (interface{}, error) {
		req := r.(int64)
		result := s.Delete(req)
		return result, nil
	}
}
