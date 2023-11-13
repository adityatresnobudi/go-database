package queries

import (
	"database/sql"
	"fmt"
)

func DeleteProduct(db *sql.DB, id int, deleteType string) (int, error) {
	sql := ""
	switch deleteType {
	case "soft", "Soft", "SOFT", "s", "S":
		sql += `UPDATE products_demo
				SET deleted_at = NOW()
				WHERE id = $1`
	case "hard", "Hard", "HARD", "h", "H":
		sql += `DELETE FROM products_demo 
				WHERE id = $1`
	default:
		return -1, fmt.Errorf("deleteType not exist")
	}

	_, err := db.Exec(sql, id)
	if err != nil {
		return -1, err
	}

	return id, nil
}
