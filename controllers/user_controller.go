package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Create user


func CreateUser(c *gin.Context) {
	var user User
	c.BindJSON(&user)
	fmt.Println(user)
	c.JSON(200, gin.H{
		"message": "User created successfully",
	})
}


// Get user

// Get all users
func GetAllUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get all users",
	})
}

// Update user

// Delete user

