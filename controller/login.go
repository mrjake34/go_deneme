package controller

import (
	"context"
	"encoding/json"
	"go_deneme/config"
	"go_deneme/models"
	"go_deneme/service"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	c := s.client.Database(config.MONGO_DB).Collection(config.MONGO_USER_COLLECTION)
	user := models.User{}
	login := models.Login{}
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	if login.Email == "" || login.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(config.ALL_FIELDS_REQUIRED))
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	filter := bson.M{config.LOGIN_EMAIL: login.Email}
	err = c.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	err = models.CheckPasswordLogin(&user, &login)
	if err != nil {
		w.Write([]byte(config.LOGIN_FAILED))
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	claims := jwt.MapClaims{
		config.SUB: user.UserID,
		config.EXP: time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	JwtKey := service.GetJwtKey()
	ss, err := token.SignedString([]byte(JwtKey))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	user.Token = ss
	update := bson.M{config.SET: bson.M{config.TOKEN: ss}}
	_, err = c.UpdateOne(context.Background(), filter, update)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(config.LOGIN_SUCCESSFUL))
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
