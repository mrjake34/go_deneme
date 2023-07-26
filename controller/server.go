package controller

import "go.mongodb.org/mongo-driver/mongo"

type Server struct {
	client *mongo.Client
}

func MongoServer(client *mongo.Client) *Server {
	return &Server{client: client}
}
