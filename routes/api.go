package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-monsters/sulivan/app/http/controllers"
)

func ApiRoutes(router *gin.Engine) {
	router.GET("/ping", controllers.Ping)
}
