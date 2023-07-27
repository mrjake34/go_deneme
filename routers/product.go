package routers

import (
	"go_deneme/controller"
	"go_deneme/service"
	"github.com/gorilla/mux"
)

func ProductRouter(r *mux.Router) {
	client := service.GetSession()
	server := controller.MongoServer(client)
	r.HandleFunc("/products", server.GetProduct).Methods("GET")
	r.HandleFunc("/products", server.SetProduct).Methods("POST")
	r.HandleFunc("/products/{productID}", server.DeleteProduct).Methods("DELETE")
	r.HandleFunc("/products/{productID}", server.UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{productID}", server.GetProductById).Methods("GET")
}