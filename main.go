package main

import (
	"fmt"
	"sullivan/bootstrap"
	"sullivan/config"
)

func main() {
	bootstrap.Start()
	//do other code
	fmt.Println(config.App.Name)
	bootstrap.Router.Run()
}