package entity

import (
	"database/sql"

	"github.com/shopspring/decimal"
)

type Product struct {
	Id    int64
	Name  sql.NullString
	Price *decimal.Decimal
	Stock sql.NullInt64
}
