package storage

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Person struct {
	Id          primitive.ObjectID `json:"_id" bson:"_id"`
	IdN         int                `json:"id" bson:"id"`
	Name        string             `json:"name" bson:"name"`
	FullName    string             `json:"fullName" bson:"fullName"`
	Email       string             `json:"email" bson:"email"`
	Gender      string             `json:"gender" bson:"gender"`
	Status      string             `json:"status" bson:"status"`
	Age         int                `json:"age" bson:"age"`
	HasChildren bool               `json:"hasChildren"`
	Engaged     bool               `json:"engaged"`
	CreatedAt   time.Time          `json:"createdAt"`
	CountryId   string             `json:"countryId" bson:"countryId"`
}

type PersonUI struct {
	Name string `json:"name"`
}
