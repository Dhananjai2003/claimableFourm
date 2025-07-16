package routes

import (
	"claimable-forum/handlers"
	"github.com/gin-gonic/gin"
	"claimable-forum/utils"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/signup", handlers.SignUp)
	r.POST("/login", handlers.Login)

	protected := r.Group("/")
protected.Use(utils.JWTauthMiddleWare())
	{
		protected.POST("/posts", handlers.CreatePost)
		protected.POST("/claim/:id",handlers.ClaimPost)
		protected.POST("/comment",handlers.CreateComment)
	}

}
