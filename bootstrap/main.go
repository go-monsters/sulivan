package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/go-monsters/sulivan/config"
	"github.com/go-monsters/sulivan/routes"
	"io/ioutil"
	"log"
)

var Router *gin.Engine

func Start() {
	//load config from env file
	config.LoadConfig()
	//create new gin
	Router = gin.Default()
	Router.Use(gin.Logger())
	Router.Use(gin.Recovery())
	registerAssets()

	//register routes
	//if production go run in production
	//if debug -> turn on debug
	routes.Register(Router)
	//logger
	// 		log on file or any other platforms. log daily
	//middleware
	// exception handler
	//response handler
	// request response logger
	//use uuid for uniq request
	//formatter for error and json
	//auth
	//crypt with key in env
	//cors
	//CSRF
	//db
	//		use GORM and postgres. use any database and multi connection. user select default connection. support multi connection
	//load template
	//add webpack vue bootstrap and ...
}

func registerAssets() {
	files, err := ioutil.ReadDir("./public")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		Router.Static("/"+f.Name(), "./public/"+f.Name())
	}
}
