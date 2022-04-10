package repos

import (
	"context"

	"git.wndv.co/go/srv/errorx"
	"git.wndv.co/sharp/app/models"
	"go.mongodb.org/mongo-driver/bson"
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

func (e *CampaignRepository) Get(ctx context.Context, filter interface{}) (models.CampaignView, error) {
	var b bson.M
	raw := e.col.FindOne(ctx, filter)
	if raw == nil {
		return models.CampaignView{}, errorx.ErrInternalServerError.Msg("Error on getting campaign")
	}

	if err := raw.Decode(&b); err != nil {
		return models.CampaignView{}, err
	}

	result, err := models.ToCampaignView(b)
	if err != nil {
		return models.CampaignView{}, err
	}

	return *result, nil
}
