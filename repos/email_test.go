package repos_test

import (
	"context"
	"fmt"
	"testing"

	"git.wndv.co/sharp/app/models"
	"git.wndv.co/sharp/app/repos"
)

func TestEmailRepository(t *testing.T) {
	e := repos.Init().Email

	t.Run("Insert case", func(t *testing.T) {
		_, err := e.Insert(context.TODO(), models.Email{
			Email:     "t4@example.com",
			Firstname: "t4",
			Lastname:  "t4",
		})
		if err != nil {
			t.Errorf("err %v", err)
		}
	})

	t.Run("Get All case", func(t *testing.T) {
		r, err := e.GetAll(context.TODO())
		fmt.Printf("r %v", r)
		if err != nil {
			t.Errorf("err %v", err)
		}
	})

}
