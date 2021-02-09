package repository

import (
	"database/sql"

	"Tienda/entities"
)

type ProductRepository struct {
	Db *sql.DB
}

func NewProductRepository(pDb *sql.DB) ProductRepository {
	return ProductRepository{
		Db: pDb,
	}
}

func (pr ProductRepository) FindAll() ([]entities.Product, error) {
	rows, err := pr.Db.Query("select * from product")

	if err != nil {
		return nil, err
	} else {
		var products []entities.Product
		for rows.Next() {
			var producto entities.Product
			err2 := rows.Scan(&producto.Id, &producto.Name, &producto.Price, &producto.Quantity, &producto.Status)
			if err2 != nil {
				return nil, err2
			} else {
				products = append(products, producto)
			}
		}
		return products, nil
	}
}

func (pl ProductRepository) Create(pro *entities.Product) bool {

	result, err2 := pl.Db.Exec("insert into product(name,price,quantity,status) values(?,?,?,?)", pro.Name, pro.Price, pro.Quantity, pro.Status)
	if err2 != nil {
		return false
	} else {
		rows, _ := result.RowsAffected()
		return rows > 0
	}
}

func (pr ProductRepository) Delete(id int64) bool {

	result, err2 := pr.Db.Exec("delete from product where id = ?", id)
	if err2 != nil {
		return false
	} else {
		rows, _ := result.RowsAffected()
		return rows > 0
	}
}
