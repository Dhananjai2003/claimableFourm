package handlers

import (
	"net/http"
	"strconv"
	"time"
	"log"
	"github.com/gin-gonic/gin"
	"claimable-forum/db"
)

func ReactComment(c *gin.Context) {
	userIDAny, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userID := userIDAny.(int)

	commentIDStr := c.Param("id")
	commentID, err := strconv.Atoi(commentIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
		return
	}

	reaction := c.Param("reaction")
	if reaction != "updoot" && reaction != "downdoot" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reaction type"})
		return
	}

	_, err = db.DB.Exec(`
		INSERT INTO comment_reactions (user_id, comment_id, reaction_type, created_at)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (user_id, comment_id) DO UPDATE SET reaction_type = EXCLUDED.reaction_type, created_at = EXCLUDED.created_at
	`, userID, commentID, reaction, time.Now())

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save reaction"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Reaction saved"})
}