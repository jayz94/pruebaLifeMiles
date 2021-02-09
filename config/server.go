package config

import (
	"Tienda/entities"
	"context"
	"encoding/json"
	"net/http"

	trans "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPHandler(endpoints Endpoints) http.Handler {
	createProductHandler := trans.NewServer(endpoints.CreateProduct, DecCreateProduct, EncCreateProduct)
	findAllProductHandler := trans.NewServer(endpoints.GetAllProduct, DecFindAllProduct, EncFindAllProduct)
	r := mux.NewRouter()
	r.Handle("/Tienda/Create", createProductHandler).Methods(http.MethodPost)
	r.Handle("/Tienda/findAll", findAllProductHandler).Methods(http.MethodGet)
	return r
}

func DecCreateProduct(_ context.Context, r *http.Request) (interface{}, error) {
	var req entities.Product
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func EncCreateProduct(ctx context.Context, writer http.ResponseWriter, response interface{}) error {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	err := json.NewEncoder(writer).Encode(response) //cast para devolver estados return json.NewEncoder(writer).Encode(response)
	if err != nil {
		http.Error(writer, err.Error(), 500)
		return err
	}
	return err
}

func DecFindAllProduct(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func EncFindAllProduct(ctx context.Context, writer http.ResponseWriter, response interface{}) error {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	err := json.NewEncoder(writer).Encode(response)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("500 - Something bad happened!"))
		return err
	}
	return err
}

func commonMiddleware(n http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		n.ServeHTTP(w, r)
	})
}
