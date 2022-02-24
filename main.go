package main

import (
	"net/http"

	"dcf-finance.com/v1/api/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	router.GET("/financial/:ticker/:function", handlers.GetFinancial)
	router.GET("/dcf/:ticker", handlers.GetDCF)

	router.Run("localhost:8080")
}
