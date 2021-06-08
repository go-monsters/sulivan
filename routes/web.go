package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-monsters/sulivan/app/http/controllers"
	"github.com/go-monsters/sulivan/app/http/controllers/auth"
)

func WebRoutes(r *gin.Engine) {
	r.GET("/", controllers.Welcome)
	r.GET("/login", auth.GetLogin)
	r.GET("/register", auth.GetRegister)
	r.POST("/login", auth.Login)
	r.POST("/register", auth.Register)
}

