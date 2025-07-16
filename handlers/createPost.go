package handlers

import ("claimable-forum/db"
	"claimable-forum/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin")

func CreatePost(c *gin.Context) {

	var req models.CreatePostRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	var postID int
	err := db.DB.QueryRow(`
		INSERT INTO posts (title, content, user_id, is_anonymous, created_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`, req.Title, req.Content, req.UserID, req.IsAnonymous, time.Now()).Scan(&postID)

	if err != nil {
		log.Println("Post insert error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	if len(req.MediaURLs) > 0 {
		stmt, err := db.DB.Prepare(`INSERT INTO media (post_id, file_url, file_type, created_at) VALUES ($1, $2, $3, $4)`)
		if err != nil {
			log.Println("Media prepare error:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to attach media"})
			return
		}
		defer stmt.Close()

		for i := range req.MediaURLs {
			fileURL := req.MediaURLs[i]
			fileType := "other"
			if i < len(req.FileTypes) {
				fileType = req.FileTypes[i]
			}
			_, err = stmt.Exec(postID, fileURL, fileType, time.Now())
			if err != nil {
				log.Println("Insert media error:", err)
			}
		}
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Post created successfully",
		"post_id": postID,
	})
}


