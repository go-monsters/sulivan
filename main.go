package main

import (
	"fmt"
	"github.com/go-monsters/sulivan/bootstrap"
	"github.com/go-monsters/sulivan/config"
)

func main() {
	bootstrap.Start()
	//do other code
	fmt.Println(config.App.Name)
	bootstrap.Router.Run()
}