package routes

import (
	"github.com/gin-gonic/gin"
	"sullivan/app/http/controllers"
)

func ApiRoute(router *gin.Engine) {
	router.GET("/ping", controllers.Ping)
}
