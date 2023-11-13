package queries

import (
	"database/sql"

	"github.com/adityatresnobudi/go-database/entity"
	"github.com/shopspring/decimal"
)

func CreateProduct(db *sql.DB, name string, stock int, price decimal.Decimal) (*entity.Product, error) {

	var product entity.Product

	sql := `INSERT INTO products_demo (name, stock, price, created_at, updated_at)
			VALUES
			($1, $2, $3, NOW(), NOW())
			`

	res, err := db.Exec(sql, name, stock, price)
	if err != nil {
		return nil, err
	}

	product.Id, err = res.LastInsertId()
	if err != nil {
		return nil, err
	}
	product.Name.String = name
	product.Stock.Int64 = int64(stock)
	product.Price = &price

	return &product, nil
}
