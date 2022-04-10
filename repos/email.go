package repos

import (
	"context"

	"git.wndv.co/sharp/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmailRepository struct {
	col *mongo.Collection
}

func (e *EmailRepository) Insert(ctx context.Context, m models.Email) (interface{}, error) {
	r, err := e.col.InsertOne(ctx, m)
	if err != nil {
		return nil, err
	}
	return r.InsertedID, nil
}

func (e *EmailRepository) GetAll(ctx context.Context) ([]models.EmailView, error) {
	var emails []models.EmailView
	cursor, err := e.col.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var email bson.M
		if err = cursor.Decode(&email); err != nil {
			return nil, err
		}

		m, err := models.ToEmailView(email)
		emails = append(emails, *m)
		if err != nil {
			return nil, err
		}
	}

	return emails, nil
}
