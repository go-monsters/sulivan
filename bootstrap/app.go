package bootstrap

import (
	"github.com/gin-gonic/gin"
	"sullivan/config"
)

func init()  {
	config.LoadConfig()
}


func Start() *gin.Engine {

	//create new gin
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
