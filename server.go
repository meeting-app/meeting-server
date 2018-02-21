package main

import (
	"github.com/ezradiniz/meeting-server/database"
	"github.com/ezradiniz/meeting-server/models"
	"github.com/ezradiniz/meeting-server/router"
)

func main() {

	e := router.New()
	database.Init()
	models.AutoMigrate()
	e.Logger.Fatal(e.Start(":8000"))
}
