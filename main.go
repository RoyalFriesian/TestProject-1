package main

import (
	"github.com/RoyalFriesian/TestProject-1/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}
