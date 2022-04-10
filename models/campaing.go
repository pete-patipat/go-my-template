package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Campaign struct {
	ID           *primitive.ObjectID `json:"ID" bson:"_id,omitempty"`
	Subject      string              `bson:"subject"`
	BodyTemplate string              `bson:"bodytemplate"`
}

type CampaignView struct {
	ID           string `json:"id" bson:"_id"`
	Subject      string
	BodyTemplate string
}
