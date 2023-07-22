package main

import (
	"context"
	"fmt"
	"go_deneme/handler"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	CONNECTION_HOST = "localhost"
	CONNECTION_PORT = ":8080"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	client := getSeesion()
	// defer client.Disconnect(context.Background())
	fmt.Println("Connected to MongoDB!")
	r := mux.NewRouter()
	server := handler.MongoServer(client)

	r.HandleFunc("/products", server.GetProduct).Methods("GET")
	r.HandleFunc("/products", server.SetProduct).Methods("POST")
	r.HandleFunc("/products/{id}", server.DeleteProduct).Methods("DELETE")
	r.HandleFunc("/products/{id}", server.UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", server.GetProductById).Methods("GET")

	http.ListenAndServe(":8080", r)
}


func getSeesion() *mongo.Client {
	var session, err = mongo.Connect(context.Background(), options.Client().ApplyURI(getEnv("MONGODB_URL", "mongodb://localhost:27017")))
	if err != nil {
		panic(err)
	}
	return session
}

func getEnv(key, defaulValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaulValue
	}
	return value
}
