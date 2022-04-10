package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Email struct {
	ID        *primitive.ObjectID `json:"ID" bson:"_id,omitempty"`
	Email     string              `bson:"email"`
	Firstname string              `bson:"firstname"`
	Lastname  string              `bson:"lastname"`
}

type EmailView struct {
	ID        string `json:"id" bson:"_id"`
	Email     string
	Firstname string
	Lastname  string
}

func ToEmailView(v interface{}) (email *EmailView, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return nil, err
	}

	err = bson.Unmarshal(data, &email)

	if err != nil {
		return nil, err
	}
	return
}
