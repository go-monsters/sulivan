package routes

import "github.com/gin-gonic/gin"

func Register(router *gin.Engine) {
	ApiRoutes(router)
	WebRoutes(router)
}
