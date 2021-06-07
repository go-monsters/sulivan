package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-monsters/sulivan/app/http/controllers"
)

func WebRoute(router *gin.Engine) {
	router.GET("/", controllers.Welcome)
}
