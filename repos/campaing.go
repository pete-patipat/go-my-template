package repos

import (
	"context"

	"git.wndv.co/sharp/app/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type CampaignRepository struct {
	col *mongo.Collection
}

func (e *CampaignRepository) Insert(ctx context.Context, m models.Campaign) (interface{}, error) {
	r, err := e.col.InsertOne(ctx, m)
	if err != nil {
		return nil, err
	}
	return r.InsertedID, nil
}
