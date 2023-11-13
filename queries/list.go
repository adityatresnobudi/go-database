package queries

import (
	"database/sql"

	"github.com/adityatresnobudi/go-database/entity"
)

func ListProducts(db *sql.DB) ([]entity.Product, error) {
	res := []entity.Product{}

	sql := `SELECT 
				id,
				name,
				stock,
				price 
			FROM products_demo`

	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var product entity.Product

		err := rows.Scan(&product.Id, &product.Name, &product.Stock, &product.Price)
		if err != nil {
			return nil, err
		}

		res = append(res, product)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return res, nil
}
