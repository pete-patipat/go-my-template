package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Email struct {
	ID        primitive.ObjectID `bson:"_id"`
	Email     string             `bson:"email"`
	Firstname string             `bson:"firstname"`
	Lastname  string             `bson:"lastname"`
}

func (e Email) NewID() Email {
	e.ID = primitive.NewObjectID()
	return e
}

type EmailRequest struct {
	Email     string
	Firstname string
	Lastname  string
}
