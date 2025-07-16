package handlers

import (
	"github.com/gin-gonic/gin"
	"claimable-forum/db"
	"net/http"
)

func ClaimPost(c *gin.Context) {

	postID := c.Param("id")

	userIDAny, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userID := userIDAny.(int)

	var ownerID int
	err := db.DB.QueryRow(`SELECT user_id FROM posts WHERE id = $1`, postID).Scan(&ownerID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	if ownerID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You cannot claim this post"})
		return
	}

	_, err = db.DB.Exec(`UPDATE posts SET is_anonymous = false WHERE id = $1`,postID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not claim post"})
		return
	}	

	c.JSON(http.StatusOK, gin.H{"message": "Post claimed successfully"})

}