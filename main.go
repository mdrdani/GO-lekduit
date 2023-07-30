package main

import (
	"GO-lekduit/config"
	"GO-lekduit/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}
