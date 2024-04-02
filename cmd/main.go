package main

import (
	"log"

	"github.com/joho/godotenv"
	_ "github.com/woozie-10/api-clean-arch/docs"
	"github.com/woozie-10/api-clean-arch/pkg/app"
)

const defaultAddress = "9090"

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// @title Students REST API
// @description This is a  API for managing students.
// @version 1.0
// @host localhost:9090
// @BasePath /
func main() {
	app := app.NewApp()
	err := app.Run(defaultAddress)
	if err != nil {
		log.Fatal(err)
	}
}
