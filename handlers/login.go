package handlers

import (
	"claimable-forum/db"
	"claimable-forum/models"
	"claimable-forum/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)


func Login(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var user models.User
	err := db.DB.QueryRow(`SELECT id, password_hash FROM users WHERE username=$1`, input.Username).Scan(&user.Id, &user.Password_hash)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if input.Password_hash != user.Password_hash {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, _ := utils.GenerateJWT(user.Id)
	c.JSON(http.StatusOK, gin.H{"token": token})
}
