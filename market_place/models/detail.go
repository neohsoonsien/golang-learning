package models

type Detail struct {
	Name            string  `json:"name" bson:"name"`
	Category        string  `json:"category" bson:"category"`
	UnitPrice       float32 `json:"unitPrice" bson:"unitPrice"`
	UnitPriceWeight string  `json:"unitPriceWeight" bson:"unitPriceWeight"`
	Count           float32 `json:"count" bson:"count"`
	CountUnit       string  `json:"countUnit" bson:"countUnit"`
}
