package app

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/woozie-10/api-clean-arch/internal/rest"
	mongo2 "github.com/woozie-10/api-clean-arch/internal/repository/mongo"
	"github.com/woozie-10/api-clean-arch/student"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type App struct {
	s student.Service
	g *gin.Engine
}

func NewApp() *App {
	db := initDB()
	studentRepo := mongo2.NewStudentRepository(db, os.Getenv("COLLECTION_NAME"))
	svc := student.NewService(studentRepo)
	ginEngine := gin.Default()
	return &App{
		s: *svc,
		g: ginEngine,
	}
}


func (a *App) Run(port string) error {
	rest.NewStudentHandler(a.g, &a.s)
	a.g.Use(
		gin.Recovery(),
		gin.Logger(),
	)
	err := a.g.Run("0.0.0.0:"+port)
	if err != nil {
		return err
	}
	return nil
}

func initDB() *mongo.Database {
	dbName := os.Getenv("DATABASE_NAME")
	connURI := os.Getenv("CONNECTION_URI")
	clientOptions := options.Client().ApplyURI(connURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Error occured while establishing connection to mongoDB")
	}
	db := client.Database(dbName)
	fmt.Println("You successfully connected to MongoDB!")
	return db
}
