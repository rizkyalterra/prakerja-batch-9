package main

import (
	"km5go/configs"
	"km5go/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main(){
	// loadEnv()
	configs.InitDatabase()
	e := echo.New()
	routes.InitRoute(e)
	e.Start(getPort())
}

func getPort() string {
	port := os.Getenv("PORT")
    if port == "" {
        port = "8000"
    }
	return ":"+port
}

func loadEnv(){
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
}



