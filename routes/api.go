package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-monsters/sulivan/app/http/controllers"
)

func ApiRoute(router *gin.Engine) {
	router.GET("/ping", controllers.Ping)
}
