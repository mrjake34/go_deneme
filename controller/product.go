package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"go_deneme/config"
	"go_deneme/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *Server) GetProduct(w http.ResponseWriter, r *http.Request) {
	c := s.client.Database(config.MONGO_DB).Collection(config.MONGO_PRODUCT_COLLECTION)
	cursor, err := c.Find(context.Background(), bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		fmt.Println(err.Error())
		return
	}
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
	c := s.client.Database(config.MONGO_DB).Collection(config.MONGO_PRODUCT_COLLECTION)
	cursor, err := c.Find(context.Background(), bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		fmt.Println(err.Error())
		return
	}
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
		w.Write([]byte(config.ALL_FIELDS_REQUIRED))
		return
	}
	_, err = c.InsertOne(context.Background(), product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(config.PRODUCT_CREATED))
}

func (s *Server) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	c := s.client.Database(config.MONGO_DB).Collection(config.MONGO_PRODUCT_COLLECTION)
	vars := mux.Vars(r)
	id := vars[config.PRODUCT_ID]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(config.ID_REQUIRED))
		return
	}
	intVal, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	filter := bson.M{config.PRODUCT_ID: intVal}
	_, err = c.DeleteOne(context.Background(), filter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(config.PRODUCT_DELETED))
}

func (s *Server) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	c := s.client.Database(config.MONGO_DB).Collection(config.MONGO_PRODUCT_COLLECTION)
	vars := mux.Vars(r)
	id := vars[config.PRODUCT_ID]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(config.ID_REQUIRED))
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
	filter := bson.M{config.PRODUCT_ID: intVal}
	update := bson.M{config.SET: bson.M{config.PRODUCT_NAME: product.Name, config.PRODUCT_COUNT: product.Count, config.PRODUCT_DESC: product.Desc, config.PRODUCT_PRICE: product.Price}}
	_, err = c.UpdateOne(context.Background(), filter, update)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(config.PRODUCT_UPDATED))
}

func (s *Server) GetProductById(w http.ResponseWriter, r *http.Request) {
	c := s.client.Database(config.MONGO_DB).Collection(config.MONGO_PRODUCT_COLLECTION)
	vars := mux.Vars(r)
	id := vars[config.PRODUCT_ID]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(config.ID_REQUIRED))
		return
	}
	intVal, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	filter := bson.M{config.PRODUCT_ID: intVal}
	product := models.Product{}
	err = c.FindOne(context.Background(), filter).Decode(&product)
	if err != nil {
		if err.Error() == config.MONGO_NO_DOCUMENTS_IN_RESULT {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(config.PRODUCT_NOT_FOUND))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set(config.CONTENT_TYPE, config.APPLICATION_JSON)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}
