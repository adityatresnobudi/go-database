package main

import (
	"fmt"
	"log"

	"github.com/adityatresnobudi/go-database/database"
	"github.com/adityatresnobudi/go-database/queries"
	"github.com/shopspring/decimal"
)

func main() {
	// connecting to DB
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("unable to connect to the database: %v", err)
	}
	defer db.Close()

	products, err := queries.ListProducts(db)
	if err != nil {
		log.Fatalf("unable to list products: %v", err)
	}

	// list products
	fmt.Println("List of products:")
	fmt.Println("=================")
	for i := 0; i < len(products); i++ {
		product := products[i]
		if product.Price == nil {
			fmt.Printf("%d. %s, %p, %v unit(s)\n", (i + 1), product.Name.String, product.Price, product.Stock.Int64)
			continue
		}
		fmt.Printf("%d. %s, Rp.%s, %v unit(s)\n", (i + 1), product.Name.String, product.Price.String(), product.Stock.Int64)
	}

	// create a new product
	product, err := queries.CreateProduct(db, "lamp", 100, decimal.NewFromInt(10000))
	if err != nil {
		log.Fatalf("unable to create a new product: %v", err)
	}
	fmt.Println("Successfully created a new product:")
	fmt.Println("===================================")
	fmt.Printf("%s, Rp.%s, %d unit(s)\n", product.Name.String, product.Price.String(), product.Stock.Int64)

	// delete a product
	id, err := queries.DeleteProduct(db, 21, "soft") // soft delete
	if err != nil {
		log.Fatalf("unable to delete a product: %v", err)
	}

	fmt.Println("Successfully deleted a product:")
	fmt.Println("===================================")
	fmt.Println("Deleted data ID: ", id)

	// implement your report product here...
	// product report
	fmt.Println("Product Report:")
	fmt.Println("===============")
	queries.ReportProducts(db)
}
