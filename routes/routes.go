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
		//Post handlers
		protected.POST("/posts", handlers.CreatePost)
		protected.POST("/reactPost/:id/react/:reaction",handlers.ReactPost)
		protected.POST("/claim/:id",handlers.ClaimPost)

		//Comment handlers
		protected.POST("/comment",handlers.CreateComment)
		protected.POST("/reactComment/:id/react/:reaction",handlers.ReactComment)
	}

}
