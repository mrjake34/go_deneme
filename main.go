package main

import (
	"fmt"
	"go_deneme/routers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println("Connected to MongoDB!")
	r := mux.NewRouter()

	routers.UserRouter(r)
	routers.ProductRouter(r)

	http.ListenAndServe("localhost:8000", r)
}
