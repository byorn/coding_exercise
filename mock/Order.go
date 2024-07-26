package mock

import (
	"time"
)

type Order struct {
	OrderID       string    `json:"order_id"`
	CustomerID    string    `json:"customer_id"`
	PlacementDate time.Time `json:"placement_date"`
	Status        string    `json:"status"`
}

func InitMockData() map[string]Order {
	orderData := map[string]Order{
		"1": Order{
			OrderID:       "1",
			CustomerID:    "10001",
			PlacementDate: time.Now(),
		},
	}

	return orderData
}
