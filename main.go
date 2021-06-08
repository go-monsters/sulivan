package main

import (
	"github.com/go-monsters/sulivan/bootstrap"
	"github.com/go-monsters/sulivan/database"
)

func main() {
	bootstrap.Start()
	bootstrap.Router.Run()
	defer database.CloseDb()
}
