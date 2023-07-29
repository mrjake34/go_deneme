package routers

import (
	"go_deneme/config"
	"go_deneme/controller"
	"go_deneme/service"

	"github.com/gorilla/mux"
)

func UserRouter(r *mux.Router) {
	client := service.GetSession()
	controller := controller.MongoServer(client)
	r.HandleFunc(config.REGISTER_PATH, controller.Register).Methods(config.POST)
	r.HandleFunc(config.LOGIN_PATH, controller.Login).Methods(config.POST)
	r.HandleFunc(config.USER_ID_PATH, controller.GetUserDetails).Methods(config.GET)
}
