// main.go
package main

import (
	"database/sql"
	"example/adapters/driven/mysql"
	"net/http"

	httpadapter "example/adapters/driving/http"
	"example/core"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq" // драйвер PostgreSQL
)

func main() {
	// Driven адаптер (БД)
	connStr := "postgres://kto:123@localhost:5432/ordersdb?sslmode=disable"
	connStr = "root:1234@tcp(127.0.0.1:3306)/ordersdb?parseTime=true"
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repo := mysql.NewOrderRepo(db)

	// Ядро: подключаем driven-порт (репозиторий)
	service := core.NewOrderService(repo)

	// Driving адаптер (HTTP): подключаем driving-порт (сервис)
	handler := httpadapter.NewOrderHandler(service)

	http.HandleFunc("/orders", handler.CreateOrder)
	http.HandleFunc("/order", handler.GetOrder)
	http.ListenAndServe(":8080", nil)
}
