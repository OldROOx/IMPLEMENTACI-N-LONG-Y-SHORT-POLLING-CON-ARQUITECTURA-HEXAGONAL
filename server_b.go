package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()

	r.GET("/status", func(c *gin.Context) {
		// Simula un retraso en la respuesta
		time.Sleep(2 * time.Second)

		if rand.Intn(2) == 0 {
			c.JSON(http.StatusOK, gin.H{"status": "pending"})
		} else {
			c.JSON(http.StatusOK, gin.H{"status": "completed"})
		}
	})

	r.GET("/long-status", func(c *gin.Context) {

		time.Sleep(10 * time.Second)
		c.JSON(http.StatusOK, gin.H{"status": "completed"})
	})

	r.Run(":8081")
}
