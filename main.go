package main

import (
	"go_deneme/config"
	"go_deneme/routers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(config.ENV_ERROR)
	}
	r := mux.NewRouter()

	routers.UserRouter(r)
	routers.ProductRouter(r)

	err = http.ListenAndServe(config.LOCAL_HOST, r)
	if err != nil {
		log.Fatal(err)
	}

}
