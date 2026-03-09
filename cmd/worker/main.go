package main

import (
	"encoding/json"
	"log"

	"event-driven-order-system/internal/kafka"
	"event-driven-order-system/internal/model"
)

func main() {

	consumer := kafka.NewConsumer()

	jobs := make(chan []byte, 100)

	// создаём 10 worker goroutines
	for i := 0; i < 10; i++ {
		go worker(i, jobs)
	}

	log.Println("Worker started with 10 goroutines")

	for {
		data, err := consumer.Read()
		if err != nil {
			log.Println(err)
			continue
		}

		jobs <- data
	}
}

func worker(id int, jobs <-chan []byte) {

	for job := range jobs {

		var order model.Order

		err := json.Unmarshal(job, &order)
		if err != nil {
			log.Println(err)
			continue
		}

		log.Printf("worker %d processed order %+v\n", id, order)
	}
}
