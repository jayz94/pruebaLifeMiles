package main

import (
	"Tienda/config"
	conf "Tienda/config"
	_ "Tienda/docs"
	repo "Tienda/repository"
	service "Tienda/services"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	httpSwagger "github.com/swaggo/http-swagger"
)

//@title Empleados API
//@version 1.0
//@description API de productos
//@contact.name Juan Carlos
//@contact.url miweb.com
//@contact.email cruzdelacruz026@gmail.com
//licence.name Apache2.0
//@host localhost:90/Tienda
//@BasePath
func main() {
	dbConnection, _ := config.GetBD()
	logger := log.NewJSONLogger(os.Stderr)
	logger = log.With(logger, "time", log.DefaultTimestamp)

	repProduct := repo.NewProductRepository(dbConnection)
	service := service.NewService(repProduct, logger)

	//Se crea documentation handler
	swaggerHandler := httpSwagger.Handler(httpSwagger.URL("http://localhost:90/swagger/doc.json"))

	//se crea el servidor
	mux := http.NewServeMux()
	e := conf.Endpoints{
		GetAllProduct: conf.MakeGetAllProductEnd(service),
		CreateProduct: conf.MakeCreateProductEnd(service),
	}

	mux.Handle("/Tienda/", conf.NewHTTPHandler(e))
	mux.Handle("/swagger/", swaggerHandler)
	//Se crea el canal para el manejo de error
	errs := make(chan error, 2)
	//Se crea la go routine para el servidor
	go func() {
		logger.Log("transport", "http", "address", "90", "msg", "listening")
		errs <- http.ListenAndServe(":90", mux)
	}()
	//Se crea la go routine para cuando se termine la ejecución
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		signal.Notify(c, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	logger.Log("terminated", <-errs)
}
