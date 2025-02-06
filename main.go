package main

import (
	"APIHEX_LPySP/application"
	"APIHEX_LPySP/config"
	"APIHEX_LPySP/infrastructure"
	"APIHEX_LPySP/infrastructure/handlers"
	"APIHEX_LPySP/infrastructure/repositories"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	config.ConnectDatabase()

	userRepo := repositories.NewUserRepository()
	productRepo := repositories.NewProductRepository()

	userService := application.NewUserService(userRepo)
	productService := application.NewProductService(productRepo)

	userHandler := handlers.NewUserHandler(userService)
	productHandler := handlers.NewProductHandler(productService)

	r := infrastructure.SetupRouter(userHandler, productHandler)

	r.GET("/check-status", func(c *gin.Context) {
		for {
			resp, err := http.Get("http://localhost:8081/status")
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			if resp.StatusCode == http.StatusOK {
				var result map[string]string
				if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}
				resp.Body.Close()

				if result["status"] == "completed" {
					c.JSON(http.StatusOK, gin.H{"message": "Proceso completado"})
					return
				}
			}

			time.Sleep(2 * time.Second)
		}
	})

	r.GET("/long-check-status", func(c *gin.Context) {
		resp, err := http.Get("http://localhost:8081/long-status")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if resp.StatusCode == http.StatusOK {
			var result map[string]string
			if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			resp.Body.Close()

			c.JSON(http.StatusOK, gin.H{"message": "Proceso completado"})
		}
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
