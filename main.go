package main

import (
	database "puja_go_bioskop/db"
	"puja_go_bioskop/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()

	r := gin.Default()

	r.POST("/bioskop", handler.CreateBioskop)
	// r.GET("/bioskop", handler.GetBioskop)
	// r.PUT("/bioskop", handler.UpdateBioskop)
	// r.DELETE("/bioskop/:id", handler.DeleteBioskop)

	r.Run(":8080")
}
