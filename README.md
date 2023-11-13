# go-database
self exercise for database with go

To improve :
database/connect.go
1. make constant for database url so to run only use go run .
2. moving defer db.Close() to main.go

queries/create.go
0. fix postgresql query from table products to products_demo (adapting to my own table name)
1. rewriting the sql variable containing postgresql query (avoiding the use of fmt, instead using placeholders when inputing variables) 
and tidy up the indentation so that the query looks more readable
2. returning res when using db.Exec() and assigning res.LastInsertId() to product.Id (so the function can return the product struct)

queries/delete.go
0. fix postgresql query from table products to products_demo (adapting to my own table name)
1. changing the use of db.Query() to db.Exec() for deleting data from table
2. changing function to return id and error (id will be used for printing the last deleted id)
3. adding extra argument deleteType for choosing whether delete is 
soft delete (query changes to updating deleted_at column to now()) or hard delete (existing query)

queries/list.go
0. fix postgresql query from table products to products_demo (adapting to my own table name)
1. rewriting the sql variable containing postgresql query so that the query looks more readable

extra :
entity/product.go
1. changing class type for name, price, and stock to sql.NullString, *decimal.Decimal, and sql.NullInt64 for handling null values

queries/report.go
1. adding new function ReportProduct

main.go
1. adding defer db.Close() from connect.go
2. handling how to print list of product when there is null value
3. implementing report product
