package mock

import (
	"time"
)

type Order struct {
	OrderID       string    `json:"order_id"`
	Items         []Item    `json:"items"`
	CustomerID    string    `json:"customer_id"`
	PlacementDate time.Time `json:"placement_date"`
	Status        string    `json:"Status"`
}

func InitMockData() map[string]Order {
	orderData := map[string]Order{
		"1": Order{
			OrderID:       "1",
			CustomerID:    "10001",
			PlacementDate: time.Now(),
			Items: []Item{{
				ItemNumber: "1",
				Quantity:   1,
				Price:      12.22,
			}},
		},
	}

	return orderData
}
