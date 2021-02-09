package services

import (
	"Tienda/entities"
	r "Tienda/repository"
	"os"
	"reflect"
	"testing"

	"github.com/go-kit/kit/log"
)

func TestNewService(t *testing.T) {

	type args struct {
		pr r.Repository
		pl log.Logger
	}
	//dependencias
	logger := log.NewJSONLogger(os.Stderr)
	logger = log.With(logger, "time", log.DefaultTimestamp)

	tests := []struct {
		name string
		args args
		want service
	}{
		{
			name: "NewService",
			args: args{
				pr: r.NewProductRepositoryMock(),
				pl: logger,
			},
			want: NewService(r.NewProductRepositoryMock(), logger),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(tt.args.pr, tt.args.pl); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_FindAll(t *testing.T) {

	logger := log.NewJSONLogger(os.Stderr)
	logger = log.With(logger, "time", log.DefaultTimestamp)

	tests := []struct {
		name    string
		s       service
		want    []entities.Product
		wantErr bool
	}{
		{
			name: "testFindAll",
			s: service{
				repo:   r.NewProductRepositoryMock(),
				logger: logger,
			},
			want: []entities.Product{
				{Id: 20, Name: "Glade", Price: 10.3, Quantity: 10, Status: true},
				{Id: 20, Name: "Glade", Price: 10.3, Quantity: 10, Status: true},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.FindAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("service.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
