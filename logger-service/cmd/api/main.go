package main

import (
	"context"
	"fmt"
	"log"
	"log-service/data"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	webPort  = "80"
	rpcPort  = "5001"
	mongoUrl = "mongodb://localhost:27017"
	grpcPort = "50001"
)

var client *mongo.Client

type Config struct {
	Models data.Models
}

func main() {
	//connect to mongo
	mongoClient, err := connectToMongo()

	if err != nil {
		log.Panic(err)
	}
	client = mongoClient

	//create context with timeout for mongo disconnect
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	//disconnect from mongo when the application exits
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Println("error disconnecting from mongo:", err)
		}
	}()

	app := Config{
		Models: data.New(client),
	}

	app.serve()
}

func (app *Config) serve() {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(), // Replace with your actual handler
	}

	log.Println("Starting server on port", webPort)
	if err := srv.ListenAndServe(); err != nil {
		log.Println("Error starting server:", err)
	}
}

func connectToMongo() (*mongo.Client, error) {
	//create connection options
	clientOptions := options.Client().ApplyURI(mongoUrl)
	clientOptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})

	//connect
	c, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("error connecting to mongo")
		return nil, err
	}
	return c, nil
}
