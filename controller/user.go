package controller

import (
	"context"
	"encoding/json"
	"go_deneme/models"
	"net/http"
)

func (s *Server) Register(w http.ResponseWriter, r *http.Request) {
	c := s.client.Database("Efes").Collection("users")
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	if user.Email == "" || user.Password == "" || user.Fullname == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("All fields are required"))
		return
	}
	err = models.HashPassword(&user)
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
	w.Write([]byte("User created successfully"))
	w.WriteHeader(http.StatusCreated)
}
func (s *Server) Login(w http.ResponseWriter, r *http.Request) {

}
