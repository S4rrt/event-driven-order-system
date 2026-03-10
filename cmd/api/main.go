package main

import (
	"database/sql"
	"event-driven-order-system/internal/handler"
	"event-driven-order-system/internal/kafka"
	"event-driven-order-system/internal/repository"
	"event-driven-order-system/internal/service"
	"log"
	"net/http"
)

func main() {

	// подключение к PostgreSQL
	connStr := "postgres://postgres:postgres@localhost:5432/orders?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// repository
	orderRepo := repository.NewOrderRepository(db)

	// kafka producer
	producer := kafka.NewProducer()

	// service
	orderService := service.NewOrderService(orderRepo, producer)

	// handler
	orderHandler := handler.NewOrderHandler(orderService)

	// routes
	http.HandleFunc("/orders", orderHandler.CreateOrder)

	log.Println("API started on :8080")

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
