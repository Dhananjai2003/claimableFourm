package routes

import (
	"claimable-forum/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/signup", handlers.SignUp)
	r.POST("/login", handlers.Login)
	r.POST("/posts", handlers.CreatePost)
}
