package main

import (
	"km5go/configs"
	"km5go/routes"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main(){
	loadEnv()
	configs.InitDatabase()
	e := echo.New()
	routes.InitRoute(e)
	e.Start(":8000")
}

func loadEnv(){
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
}



