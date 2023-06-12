package main

import (
	"github.com/gin-gonic/gin"
	"programmerData/controller/programmercontroller"
	"programmerData/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/api/programmers", programmercontroller.Index)
	r.GET("/api/programmer/:id", programmercontroller.Show)
	r.POST("/api/programmer", programmercontroller.Create)
	r.PUT("/api/programmer/:id", programmercontroller.Update)
	r.DELETE("/api/programmer/:id", programmercontroller.Delete)

	r.Run()
}
