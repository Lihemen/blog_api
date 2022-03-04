package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lihemen/gin_api/controllers"
)



func main() {

	router := gin.Default()

	router.POST("/users", controllers.CreateUser)
	router.GET("/users", controllers.GetAllUsers)
	
	router.Run(":8080")
}