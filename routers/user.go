package routers

import (
	"go_deneme/controller"
	"go_deneme/service"

	"github.com/gorilla/mux"
)

func UserRouter(r *mux.Router) {
	client := service.GetSession()
	controller := controller.MongoServer(client)
	r.HandleFunc("/register", controller.Register).Methods("POST")
	r.HandleFunc("/login", controller.Login).Methods("POST")
}
