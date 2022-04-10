package repos_test

import (
	"context"
	"testing"

	"git.wndv.co/sharp/app/models"
	"git.wndv.co/sharp/app/repos"
)

func TestCampaignRepository(t *testing.T) {
	e := repos.Init().Campaign

	err := e.Insert(context.TODO(), models.Campaign{
		Subject:      "t1",
		BodyTemplate: "t1",
	}.NewID())
	if err != nil {
		t.Errorf("err %v", err)
	}
}
