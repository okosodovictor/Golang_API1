package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pluralsight/inventoryservice/database"
	"github.com/pluralsight/inventoryservice/product"
)

const apiBasePath="/api"
func main() {
	database.SetupDatabase()
	product.SetupRoutes(apiBasePath)
	http.ListenAndServe(":5000", nil)
}
