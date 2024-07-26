package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	server := &Server{
		router: gin.Default(),
	}
	server.router.POST("/orders", basicAuth(createOrderController))
	server.router.PUT("/orders/:id", basicAuth(updateOrderController))
	server.router.GET("/orders", basicAuth(searchOrdersController))
	return server
}

func basicAuth(fn gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password, ok := c.Request.BasicAuth()
		if !ok || !validateCredentials(username, password) {
			c.Header("WWW-Authenticate", `Basic realm="restricted"`)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		fn(c)
	}
}

func validateCredentials(username, password string) bool {
	// Replace with your actual username and password validation logic
	return username == "admin" && password == "password"
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
