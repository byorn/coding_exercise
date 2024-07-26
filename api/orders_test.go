package api

import (
	"bytes"
	"encoding/json"
	db "kindred/mock"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/orders", basicAuth(createOrderController))
	router.PUT("/orders/:id", basicAuth(updateOrderController))
	router.GET("/orders", basicAuth(searchOrdersController))
	return router
}

func TestCreateOrderHandler(t *testing.T) {
	router := setupRouter()

	order := db.Order{
		CustomerID: "12345",
		Items: []db.Item{
			{ItemNumber: "item-001", Quantity: 2, Price: 19.99},
			{ItemNumber: "item-002", Quantity: 1, Price: 9.99},
		},
	}
	orderJSON, _ := json.Marshal(order)

	req, _ := http.NewRequest("POST", "/orders", bytes.NewBuffer(orderJSON))
	req.SetBasicAuth("admin", "password")
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var createdOrder db.Order
	json.Unmarshal(w.Body.Bytes(), &createdOrder)

	assert.NotEmpty(t, createdOrder.OrderID)
	assert.Equal(t, "12345", createdOrder.CustomerID)
	assert.Equal(t, 2, len(createdOrder.Items))
	assert.Equal(t, "Pending", createdOrder.Status)
	assert.WithinDuration(t, time.Now(), createdOrder.PlacementDate, time.Second)
}

func TestUpdateOrderHandler(t *testing.T) {
	router := setupRouter()

	// First, create an order to update
	order := db.Order{
		CustomerID: "12345",
		Items: []db.Item{
			{ItemNumber: "item-001", Quantity: 2, Price: 19.99},
			{ItemNumber: "item-002", Quantity: 1, Price: 9.99},
		},
	}
	orderID := generateOrderID()
	order.OrderID = orderID
	order.PlacementDate = time.Now()
	orders[orderID] = order

	// Now, update the order
	updatedOrder := db.Order{
		CustomerID: "12345",
		Items: []db.Item{
			{ItemNumber: "item-001", Quantity: 3, Price: 19.99},
			{ItemNumber: "item-002", Quantity: 1, Price: 9.99},
		},
		Status: "shipped",
	}
	updatedOrderJSON, _ := json.Marshal(updatedOrder)

	req, _ := http.NewRequest("PUT", "/orders/"+orderID, bytes.NewBuffer(updatedOrderJSON))
	req.SetBasicAuth("admin", "password")
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var returnedOrder db.Order
	json.Unmarshal(w.Body.Bytes(), &returnedOrder)

	assert.Equal(t, orderID, returnedOrder.OrderID)
	assert.Equal(t, "12345", returnedOrder.CustomerID)
	assert.Equal(t, 2, len(returnedOrder.Items))
	assert.Equal(t, "shipped", returnedOrder.Status)
}

func TestSearchOrdersHandler(t *testing.T) {
	router := setupRouter()
	orders := map[string]db.Order{}
	// Create a couple of orders for testing
	order1 := db.Order{
		OrderID:       generateOrderID(),
		CustomerID:    "12345",
		PlacementDate: time.Now(),
		Items: []db.Item{
			{ItemNumber: "item-001", Quantity: 2, Price: 19.99},
			{ItemNumber: "item-002", Quantity: 1, Price: 9.99},
		},
		Status: "pending",
	}
	orders[order1.OrderID] = order1

	order2 := db.Order{
		OrderID:       generateOrderID(),
		CustomerID:    "12345",
		PlacementDate: time.Now(),
		Items: []db.Item{
			{ItemNumber: "item-003", Quantity: 5, Price: 29.99},
		},
		Status: "shipped",
	}

	orders[order2.OrderID] = order2

	req, _ := http.NewRequest("GET", "/orders?customer_id=12345", nil)
	req.SetBasicAuth("admin", "password")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var result []db.Order
	json.Unmarshal(w.Body.Bytes(), &result)

	assert.Equal(t, 2, len(result))

}
