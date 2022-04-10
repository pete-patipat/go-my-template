package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Campaign struct {
	ID           primitive.ObjectID `bson:"_id"`
	Subject      string             `bson:"subject"`
	BodyTemplate string             `bson:"bodytemplate"`
}

func (c Campaign) NewID() Campaign {
	c.ID = primitive.NewObjectID()
	return c
}
