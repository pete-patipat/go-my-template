package repos

import (
	"context"

	"git.wndv.co/sharp/app/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmailRepository struct {
	col *mongo.Collection
}

func (e *EmailRepository) Insert(ctx context.Context, m models.Email) error {
	_, err := e.col.InsertOne(ctx, m)
	if err != nil {
		return err
	}
	return nil
}
