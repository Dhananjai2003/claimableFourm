package handlers

import (
	"claimable-forum/db"
	"claimable-forum/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)


func SignUp(c *gin.Context) {

	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	_, err := db.DB.Exec(`INSERT INTO users (username, password_hash) VALUES ($1, $2)`, input.Username, input.Password_hash)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created"})

}
