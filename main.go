package main

import (
	"github.com/AryaWgna/go-restapi-gin/controllers/productcontroller"
	"github.com/AryaWgna/go-restapi-gin/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/mahasiswa", productcontroller.Index)
	r.GET("/api/Mahasiswa/:id", productcontroller.Show)
	r.POST("/api/Mahasiswa", productcontroller.Create)
	r.PUT("/api/Mahasiswa/:id", productcontroller.Update)
	r.DELETE("/api/Mahasiswa", productcontroller.Delete)

	r.Run()
}
