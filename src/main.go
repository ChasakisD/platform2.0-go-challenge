package main

import (
	"gwi/assignment/app"
)

// @title GWI Go Challenge API
// @version 1.0
// @description This is a programming assignment for GWI Go Backend position

// @contact.name Dionysis Chasakis
// @contact.email chasakisd@gmail.com

// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	app.Run()
}
