package service

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetSession() *mongo.Client {
	session, err := mongo.Connect(context.Background(), options.Client().ApplyURI(getEnv("MONGODB_URL", "mongodb://localhost:5554")))
	if err != nil {
		panic(err)
	}
	start := time.Now()
	for {
		err = session.Ping(context.Background(), nil)
		if err == nil {
			break
		}
		if time.Since(start) > 10*time.Second {
			panic(err)
		}
		time.Sleep(100 * time.Millisecond)
	}
	return session
}

func getEnv(key, defaulValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaulValue
	}
	return value
}

