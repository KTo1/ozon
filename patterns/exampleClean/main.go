package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql" // пример драйвера

	"exampleClean/adapters"
	"exampleClean/usecases"
)

func main() {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		connStr = "root:1234@tcp(127.0.0.1:3306)/ordersdb?parseTime=true"
	}

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("DB ping failed:", err)
	}

	// Сборка цепочки зависимостей
	repo := adapters.NewSQLOrderRepository(db)            // выходной адаптер (Gateway)
	interactor := usecases.NewOrderInteractor(repo)       // интерактор (Use Case)
	controller := adapters.NewOrderController(interactor) // входной адаптер (Controller)

	http.HandleFunc("/orders", controller.CreateOrder)
	http.HandleFunc("/order", controller.GetOrder)

	log.Println("Сервер на :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
