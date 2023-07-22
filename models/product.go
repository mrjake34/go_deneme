package models

var products = []Product{
	Product{ID: "1", Name: "Macbook Pro", Count: 10, Desc: "Macbook Pro 2019", Price: 1000},
	Product{ID: "2", Name: "Macbook Air", Count: 20, Desc: "Macbook Air 2019", Price: 800},
	Product{ID: "3", Name: "Macbook", Count: 30, Desc: "Macbook 2019", Price: 600},
	Product{ID: "4", Name: "Macbook Pro", Count: 10, Desc: "Macbook Pro 2019", Price: 1000},
	Product{ID: "5", Name: "Macbook Air", Count: 20, Desc: "Macbook Air 2019", Price: 800},
	Product{ID: "6", Name: "Macbook", Count: 30, Desc: "Macbook 2019", Price: 600},
	Product{ID: "7", Name: "Macbook Pro", Count: 10, Desc: "Macbook Pro 2019", Price: 1000},
	Product{ID: "8", Name: "Macbook Air", Count: 20, Desc: "Macbook Air 2019", Price: 800},
	Product{ID: "9", Name: "Macbook", Count: 30, Desc: "Macbook 2019", Price: 600},
	Product{ID: "10", Name: "Macbook Pro", Count: 10, Desc: "Macbook Pro 2019", Price: 1000},
}

type Product struct {
	ID    string `json:"id" bson:"_id"`
	Name  string `json:"name" bson:"name"`
	Count int    `json:"count" bson:"count"`
	Desc  string `json:"desc" bson:"desc"`
	Price int    `json:"price" bson:"price"`
}
