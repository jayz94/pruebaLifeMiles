package repository

/*
func TestProductRepository_FindAll(t *testing.T) {
	tests := []struct {
		name    string
		pr      ProductRepository
		want    []entities.Product
		wantErr bool
	}{
		{
			name: "testFindAll",
			pr:   NewProductRepositoryMock(),
			want: []entities.Product{
				{Id: 20, Name: "Glade", Price: 10.3, Quantity: 10, Status: true},
				{Id: 20, Name: "Glade", Price: 10.3, Quantity: 10, Status: true},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.pr.FindAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductRepository.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductRepository.FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
*/
