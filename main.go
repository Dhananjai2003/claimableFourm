package main 

import (
	"github.com/gin-gonic/gin"
	"claimable-forum/routes"
	"claimable-forum/db"
)

func main() {

	db.Connect()
	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":8080")

}