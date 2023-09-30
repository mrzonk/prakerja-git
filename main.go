package main

import (
	//"os"
	"prakerja11/configs"
	"prakerja11/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)


func main(){
	// loadEnv()
	configs.InitDatabase()
	e := echo.New()
	routes.InitRoutes(e)
	e.Start(":3436")
}

func loadEnv(){
	err := godotenv.Load()
  	if err != nil {
    	panic("Error loading .env file")
  	}
}


