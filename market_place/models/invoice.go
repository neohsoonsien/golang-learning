package models

type Invoice struct {
	PrincipalId string   `json:"principalId" bson:"principalId"`
	Number      string   `json:"number" bson:"number"`
	Vendor      string   `json:"vendor" bson:"vendor"`
	DateTime    int32    `json:"dateTime" bson:"dateTime"`
	Details     []Detail `json:"details" bson:"details"`
}
