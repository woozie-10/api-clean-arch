package main

import (
	"log"
	"os"

	_ "github.com/woozie-10/api-clean-arch/docs"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/woozie-10/api-clean-arch/internal/database"
	"github.com/woozie-10/api-clean-arch/internal/repository/mongo"
	"github.com/woozie-10/api-clean-arch/internal/rest"
	"github.com/woozie-10/api-clean-arch/student"
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
	dbName := os.Getenv("DATABASE_NAME")
	collectionName := os.Getenv("COLLECTION_NAME")
	connURI := os.Getenv("CONNECTION_URI")
	db, err := database.InitDB(connURI, dbName)
	if err != nil {
		log.Fatal("failed to open connection to database", err)
	}
	g := gin.Default()
	studentRepo := mongo.NewStudentRepository(db, collectionName)
	svc := student.NewService(studentRepo)
	rest.NewStudentHandler(g, svc)
	address := os.Getenv("SERVER_ADDRESS")
	if address == "" {
		address = defaultAddress
	}
	if err := g.Run("0.0.0.0:" + address); err != nil {
		log.Fatal("Server startup error")
	}
}
