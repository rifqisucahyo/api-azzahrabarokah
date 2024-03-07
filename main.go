package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Membuat instance Gin
	r := gin.Default()

	// Menyajikan halaman HTML
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello World"})
	})

	// Menjalankan server pada port 8080
	r.Run(":8080")
}
