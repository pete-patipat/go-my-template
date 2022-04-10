package repos_test

import (
	"context"
	"fmt"
	"testing"

	"git.wndv.co/sharp/app/models"
	"git.wndv.co/sharp/app/repos"
	"go.mongodb.org/mongo-driver/bson"
)

func TestCampaignRepository(t *testing.T) {
	var id interface{}
	e := repos.Init().Campaign
	t.Run("Insert campaign case", func(t *testing.T) {
		r, err := e.Insert(context.TODO(), models.Campaign{
			Subject:      "t1",
			BodyTemplate: "t1",
		})
		if err != nil {
			t.Errorf("err %v", err)
		}

		id = r

		fmt.Printf("result %v", r)
	})

	t.Run("Get campaign case", func(t *testing.T) {
		r, err := e.Get(context.TODO(), bson.M{"_id": id})
		if err != nil {
			t.Errorf("err %v", err)
		}

		fmt.Printf("result %v", r)
	})
}
