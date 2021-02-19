package bootstrap

import (
	"github.com/gin-gonic/gin"
	"sullivan/config"
	"sullivan/routes"
)

var Router *gin.Engine

func Start() {
	//load config from env file
	config.LoadConfig()
	//create new gin
	Router = gin.Default()
	//register routes
	routes.Register(Router)
	//logger
	//middleware
	//auth
	//cors
	//csrf
	//db
}
