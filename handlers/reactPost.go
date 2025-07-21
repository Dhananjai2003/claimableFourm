package handlers

import (
	"net/http"
	"strconv"
	"time"
	"github.com/gin-gonic/gin"
	"claimable-forum/db"
	"log"
)


func ReactPost(c *gin.Context) {

	userIDAny, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userID := userIDAny.(int)

	postIDStr := c.Param("id")
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	reaction := c.Param("reaction")
	if reaction != "updoot" && reaction != "downdoot" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reaction type"})
		return
	}

	_, err = db.DB.Exec(`
		INSERT INTO post_reactions (user_id, post_id, reaction, created_at)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (user_id, post_id) DO UPDATE SET reaction = EXCLUDED.reaction, created_at = EXCLUDED.created_at
	`, userID, postID, reaction, time.Now())

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save reaction"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Reaction saved"})

}