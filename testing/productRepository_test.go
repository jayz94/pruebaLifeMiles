package testing

import (
	"github.com/stretchr/testify/mock"
)

type mockRepository struct {
	mock.Mock
}

/*
func TestFindAll1(t *testing.T) {
	dbConnection, _ := config.GetBD()
	mockRepository := new(mockRepository)
	returnD := entities.Product{Id: 20, Name: "Glade", Price: 10.3, Quantity: 10, Status: true}
	mockRepository.On("FindAll").Return(returnD, nil)
	result, _ := mockRepository.FindAll()
	assert.Equal(t, "Glade", result[0].Name)
}

func TestFindAll(t *testing.T) {
	dbConnection, _ := config.GetBD()
	mockRepository := r.NewProductRepository(dbConnection)

	returnD := []entities.Product{{Id: 20, Name: "Glade", Price: 10.3, Quantity: 10, Status: true}, {Id: 20, Name: "Glade", Price: 10.3, Quantity: 10, Status: true}}

	mockRepository.FindAll().Return(returnD, nil)
	result, _ := mockRepository.FindAll()
	assert.Equal(t, "Glade", result[0].Name)
}
*/
