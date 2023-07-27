package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"go_deneme/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

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

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

func (s *Server) SetProduct(w http.ResponseWriter, r *http.Request) {
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
	product := models.Product{}
	if products != nil {
		product.ProductID = products[len(products)-1].ProductID + 1
	} else {
		product.ProductID = 1
	}
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if product.Name == "" || product.Desc == "" || product.Price == 0 || product.Count == 0 {
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

func (s *Server) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	c := s.client.Database("Efes").Collection("product")
	vars := mux.Vars(r)
	id := vars["productID"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID is required"))
		return
	}
	intVal, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	filter := bson.M{"productID": intVal}
	_, err = c.DeleteOne(context.Background(), filter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Product deleted successfully"))
}

func (s *Server) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	c := s.client.Database("Efes").Collection("product")
	vars := mux.Vars(r)
	id := vars["productID"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID is required"))
		return
	}
	intVal, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	product := models.Product{}
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	filter := bson.M{"productID": intVal}
	update := bson.M{"$set": bson.M{"name": product.Name, "count": product.Count, "desc": product.Desc, "price": product.Price}}
	_, err = c.UpdateOne(context.Background(), filter, update)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Product updated successfully"))
}

func (s *Server) GetProductById(w http.ResponseWriter, r *http.Request) {
	c := s.client.Database("Efes").Collection("product")
	vars := mux.Vars(r)
	id := vars["productID"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID is required"))
		return
	}
	intVal, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	filter := bson.M{"productID": intVal}
	product := models.Product{}
	err = c.FindOne(context.Background(), filter).Decode(&product)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Product not found"))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}
