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

// func (pwf *ProductFactWorker) start() error {
// 	call := pwf.client.Database("Efes").Collection("product")
// 	ticker := time.NewTicker(2 * time.Second)
// 	for {
// 		resp, err := http.Get("localhost:8080")

// 		<-ticker.C
// 	}
// }

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

	http.ListenAndServe(":8080", r)
}

// func GetProduct(w http.ResponseWriter, r *http.Request) {
// 	session, err := mongo.Connect(context.Background(), options.Client().ApplyURI(getEnv("MONGODB_URL", "mongodb://localhost:27017")))
// 	defer session.Disconnect(context.Background())

// 	if err != nil {
// 		panic(err)
// 	}

// 	c := session.Database("Efes").Collection("product")
// 	fmt.Println(c)
// 	cursor, err := c.Find(context.TODO(), nil)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte(err.Error()))
// 		fmt.Println(err.Error())
// 		return
// 	}
// 	var products []models.Product
// 	if err = cursor.All(context.TODO(), &products); err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte(err.Error()))
// 		fmt.Println(err.Error())
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(products)
// }

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
