package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"go_deneme/models"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	client *mongo.Client
}

func MongoServer(client *mongo.Client) *Server {
	return &Server{client: client}
}

func (s *Server) GetProduct(w http.ResponseWriter, r *http.Request) {
	c := s.client.Database("Efes").Collection("product")
	cursor, err := c.Find(context.Background(), bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Cursor", cursor)
	var products []models.Product
	if err = cursor.All(context.Background(), &products); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

func (s *Server) SetProduct(w http.ResponseWriter, r *http.Request) {
	c := s.client.Database("Efes").Collection("product")
	product := models.Product{}
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	existingProduct := models.Product{}
	filter := bson.M{"_id": product.ID}
	err = c.FindOne(context.Background(), filter).Decode(&existingProduct)
	if err == nil {

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Product with the same ID already exists"))
		return
	} else if err != mongo.ErrNoDocuments {

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if product.Name == "" && product.Desc == "" && product.Price == 0 && product.Count == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("All fields are required"))
		return
	}
	_, err = c.InsertOne(context.Background(), product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Product created successfully"))
}
