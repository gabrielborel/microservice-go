package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/gabrielborel/microservice-go/internal/order/infra/database"
	"github.com/gabrielborel/microservice-go/internal/order/usecases"
	"github.com/gabrielborel/microservice-go/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	db, err := sql.Open("postgres", "postgres://docker:docker@postgres:5432/orders?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	repository := database.NewOrderRepository(db)
	http.HandleFunc("/total", func(w http.ResponseWriter, r *http.Request) {
		usecase := usecases.NewGetTotalOrdersUseCase(repository)
		total, err := usecase.Execute()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(total)
	})
	go http.ListenAndServe(":8181", nil)

	usecase := usecases.NewCalculateFinalPriceUseCase(repository)
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	maxWorkers := 10
	wg := sync.WaitGroup{}

	out := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, out)

	wg.Add(maxWorkers)
	for i := 0; i < maxWorkers; i++ {
		defer wg.Done()
		go worker(out, usecase, i)
	}
	wg.Wait()
}

func worker(deliveryMessage <-chan amqp.Delivery, usecase *usecases.CalculateFinalPriceUseCase, workerId int) {
	for msg := range deliveryMessage {
		var input usecases.OrderInputDTO
		err := json.Unmarshal(msg.Body, &input)
		if err != nil {
			fmt.Println("Error while unmarshalling message", err)
		}

		input.Tax = 10.0
		_, err = usecase.Execute(input)
		if err != nil {
			panic(err)
		}

		println("Received a message from worker: %s", workerId)
		msg.Ack(false)
	}
}
