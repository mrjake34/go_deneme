package models

type Product struct {
	ObjectID  string `json:"_id,omitempty" bson:"_id,omitempty"`
	ProductID int    `json:"productID" bson:"productID"`
	Name      string `json:"name" bson:"name"`
	Count     int    `json:"count" bson:"count"`
	Desc      string `json:"desc" bson:"desc"`
	Price     int    `json:"price" bson:"price"`
}
