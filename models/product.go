package models

type Product struct {
	ID    string `json:"id" bson:"_id"`
	Name  string `json:"name" bson:"name"`
	Count int    `json:"count" bson:"count"`
	Desc  string `json:"desc" bson:"desc"`
	Price int    `json:"price" bson:"price"`
}
