package routes

import "github.com/gin-gonic/gin"

func Register(router *gin.Engine) {
	ApiRoute(router)
	WebRoute(router)
}
