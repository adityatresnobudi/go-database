package queries

import (
	"database/sql"
	"fmt"
	"log"
)

func ReportProducts(db *sql.DB) {
	var (
		product_group string
		count_product int
	)
	i := 1
	sql := `SELECT
				(CASE
					WHEN pd.price >= 10000 THEN 'Expensive Product'
					WHEN pd.price < 10000 AND pd.price > 5000 THEN 'Good Deal Product'
					ELSE 'Cheap Product'
				END) AS product_group,
				COUNT(pd.id) AS count_product
			FROM products_demo pd 
			GROUP BY product_group`

	rows, err := db.Query(sql)
	if err != nil {
		log.Fatalf("unable to create report: %v", err)
	}

	for rows.Next() {
		err := rows.Scan(&product_group, &count_product)
		if err != nil {
			log.Fatalf("unable to create report: %v", err)
		}

		fmt.Printf("%d. %s: %d product(s)\n", i, product_group, count_product)
		i++
	}

	err = rows.Err()
	if err != nil {
		log.Fatalf("unable to create report: %v", err)
	}
}
