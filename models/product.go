package models

type Product struct {
	ObjectID      string          `json:"_id,omitempty" bson:"_id,omitempty"`
	ProductID     int             `json:"productID" bson:"productID"`
	Name          string          `json:"name" bson:"name"`
	Count         int             `json:"count" bson:"count"`
	Discount      int             `json:"discount" bson:"discount"`
	Desc          string          `json:"desc" bson:"desc"`
	Category      string          `json:"category" bson:"category"`
	StandOut      string          `json:"standOut" bson:"standOut"`
	Price         float32         `json:"price" bson:"price"`
	Image         string          `json:"image" bson:"image"`
	Specification []Specification `json:"specification" bson:"specification"`
}

type Specification struct {
	Name   string   `json:"name" bson:"name"`
	Values []string `json:"values" bson:"values"`
}
