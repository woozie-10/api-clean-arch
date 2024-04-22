package app

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/woozie-10/api-clean-arch/assessment"
	mongo2 "github.com/woozie-10/api-clean-arch/internal/repository/mongo"
	"github.com/woozie-10/api-clean-arch/internal/rest"
	"github.com/woozie-10/api-clean-arch/student"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type App struct {
	sSvc student.StudentService
	aSvc assessment.AssessmentService
	g    *gin.Engine
}

func NewApp() *App {
	db := initDB()
	studentRepo := mongo2.NewStudentRepository(db, "students")
	assessmentsRepo := mongo2.NewAssessmentRepository(db, "assessments")
	sSvc := student.NewStudentService(studentRepo)
	aSvc := assessment.NewAssessmentService(assessmentsRepo)
	ginEngine := gin.Default()
	return &App{
		sSvc: *sSvc,
		aSvc: *aSvc,
		g:    ginEngine,
	}
}

func (a *App) Run(port string) error {
	rest.NewStudentHandler(a.g, &a.sSvc, &a.aSvc)
	a.g.Use(
		gin.Recovery(),
		gin.Logger(),
	)
	err := a.g.Run("0.0.0.0:" + port)
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
