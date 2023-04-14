package main

import (
	"alterra2/config"
	"alterra2/route"
	"os"
)

func main()  {
	config.InitDB()
	e := route.InitRoute()
	e.Logger.Fatal(e.Start(":" + getPort()))	
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8080"
	}
	return port
}
