package models

type Detail struct {
	Name            string  `json:"name" bson:"name"`
	Category        string  `json:"category" bson:"category"`
	UnitPrice       float32 `json:"unitPrice" bson:"unitPrice"`
	UnitPriceWeight string  `json:"unitPriceWeight" bson:"unitPriceWeight"`
	Count           int32   `json:"count" bson:"count"`
	CountUnit       int32   `json:"countUnit" bson:"countUnit"`
}
