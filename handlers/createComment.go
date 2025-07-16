package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"claimable-forum/models"
	"claimable-forum/db"
	"log"
)

func CreateComment(c *gin.Context) {

	var req models.CreateCommentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"Error":"Invalid Request"})
		return
	}

	userIDAny, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userID := userIDAny.(int)

	query := `
		INSERT INTO comments (post_id, parent_id, user_id, content, is_anonymous, created_at)
		VALUES ($1, $2, $3, $4, $5, NOW())
		RETURNING id;
	`

	var commentID int
	err := db.DB.QueryRow(query, req.PostID, req.ParentID, userID, req.Content, false).Scan(&commentID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save comment"})
		return
	}
	

	c.JSON(http.StatusCreated, gin.H{
		"message":    "Comment added",
		"comment_id": commentID,
	})

}