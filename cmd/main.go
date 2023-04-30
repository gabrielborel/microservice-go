package main

import (
	"database/sql"

	"github.com/gabrielborel/microservice-go/internal/order/infra/database"
	"github.com/gabrielborel/microservice-go/internal/order/usecases"
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
	usecase := usecases.NewCalculateFinalPriceUseCase(repository)
	_, err = usecase.Execute(usecases.OrderInputDTO{
		ID:    "1234567",
		Price: 100.0,
		Tax:   15.0,
	})
	if err != nil {
		panic(err)
	}
}
