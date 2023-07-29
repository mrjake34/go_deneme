package controller

import (
	"context"
	"encoding/json"
	"go_deneme/config"
	"go_deneme/models"
	"net/http"
)

func (s *Server) Register(w http.ResponseWriter, r *http.Request) {
	c := s.client.Database(config.MONGO_DB).Collection(config.MONGO_USER_COLLECTION)
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	if user.Email == "" || user.Password == "" || user.Fullname == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(config.ALL_FIELDS_REQUIRED))
		return
	}
	err = models.HashPasswordUser(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	_, err = c.InsertOne(context.Background(), user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(config.USER_CREATED))
	w.WriteHeader(http.StatusCreated)
}