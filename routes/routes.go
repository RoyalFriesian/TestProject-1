package routes

import (
	"github.com/RoyalFriesian/TestProject-1/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/login", controllers.Login)
	r.POST("/register", controllers.Register)
}
