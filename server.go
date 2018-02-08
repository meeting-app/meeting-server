package main

import "github.com/ezradiniz/meeting-server/router"

func main() {

	e := router.New()

	e.Logger.Fatal(e.Start(":8000"))
}
