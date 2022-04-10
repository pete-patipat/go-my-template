package repos_test

import (
	"context"
	"testing"

	"git.wndv.co/sharp/app/models"
	"git.wndv.co/sharp/app/repos"
)

func TestEmailRepository(t *testing.T) {
	e := repos.Init().Email

	err := e.Insert(context.TODO(), models.Email{
		Email:     "p1111@example.com",
		Firstname: "f22222",
		Lastname:  "l33333",
	}.NewID())
	if err != nil {
		t.Errorf("err %v", err)
	}
}
