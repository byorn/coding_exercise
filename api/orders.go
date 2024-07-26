package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	db "kindred/mock"
	"net/http"
	"time"
)

var orders = db.InitMockData()

func createOrderController(c *gin.Context) {
	var order db.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	orderID := generateOrderID()
	order.OrderID = orderID
	order.PlacementDate = time.Now()
	order.Status = "Pending"
	orders[orderID] = order
	c.JSON(http.StatusOK, order)
}

func updateOrderController(c *gin.Context) {
	orderID := c.Param("id")
	var order db.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if _, ok := orders[orderID]; ok {
		order.OrderID = orderID
		orders[orderID] = order
		c.JSON(http.StatusOK, order)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
	}
}

func searchOrdersController(c *gin.Context) {
	customerID := c.Query("customer_id")
	var result []db.Order
	for _, order := range orders {
		if order.CustomerID == customerID {
			result = append(result, order)
		}
	}
	c.JSON(http.StatusOK, result)
}

func generateOrderID() string {
	return uuid.New().String()
}
