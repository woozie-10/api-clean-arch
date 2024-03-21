package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDB(connURI, dbName string) (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(connURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}
	db := client.Database(dbName)
	fmt.Println("You successfully connected to MongoDB!")
	return db, nil
}
