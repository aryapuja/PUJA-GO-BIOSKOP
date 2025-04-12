package main

import (
	"os"
	database "puja_go_bioskop/db"
	"puja_go_bioskop/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()

	r := gin.Default()

	r.POST("/bioskop", handler.CreateBioskop)
	r.GET("/bioskop", handler.GetBioskop)
	r.GET("/bioskop/:id", handler.GetBioskopInfo)
	r.PUT("/bioskop/:id", handler.UpdateBioskop)
	r.DELETE("/bioskop/:id", handler.DeleteBioskop)

	// r.Run(":8080")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback ke 8080 kalau belum ada
	}
	r.Run(":" + port)
}
