package database

import (
	"database/sql"

	"github.com/gabrielborel/microservice-go/internal/order/entity"
	_ "github.com/lib/pq"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	stmt, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES ($1, $2, $3, $4);")
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) GetTotal() (int, error) {
	var total int
	err := r.Db.QueryRow("SELECT COUNT(*) FROM orders;").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}
