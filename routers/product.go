package routers

import (
	"go_deneme/config"
	"go_deneme/controller"
	"go_deneme/service"

	"github.com/gorilla/mux"
)

func ProductRouter(r *mux.Router) {
	client := service.GetSession()
	server := controller.MongoServer(client)
	r.HandleFunc(config.PRODUCT_PATH, server.GetProduct).Methods(config.GET)
	r.HandleFunc(config.PRODUCT_PATH, server.SetProduct).Methods(config.POST)
	r.HandleFunc(config.PRODUCT_ID_PATH, server.DeleteProduct).Methods(config.DELETE)
	r.HandleFunc(config.PRODUCT_ID_PATH, server.UpdateProduct).Methods(config.PUT)
	r.HandleFunc(config.PRODUCT_ID_PATH, server.GetProductById).Methods(config.GET)
}
