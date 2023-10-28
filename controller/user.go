package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"go_deneme/config"
	"go_deneme/models"
	"go_deneme/service"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Server) GetUserDetails(w http.ResponseWriter, r *http.Request) {
	c := s.client.Database(config.MONGO_DB).Collection(config.MONGO_USER_COLLECTION)
	vars := mux.Vars(r)
	id := vars[config.USER_ID]
	fmt.Println(id)
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(config.ID_REQUIRED))
		return
	}
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	filter := bson.M{config.USER_ID: objectID}
	user := models.User{}
	err = c.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if err.Error() == config.MONGO_NO_DOCUMENTS_IN_RESULT {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(config.USER_NOT_FOUND))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	JwtKey := service.GetJwtKey()
	secretKey := []byte(JwtKey)

	token, err := jwt.Parse(user.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok { // Check the signing method
			return nil, fmt.Errorf(config.UNEXPECTED_SIGNING_METHOD, token.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims[config.SUB], claims[config.EXP])
		w.Header().Set(config.CONTENT_TYPE, config.APPLICATION_JSON)
		json.NewEncoder(w).Encode(user)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	
}
