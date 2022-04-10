package web

import (
	"git.wndv.co/go/srv/gin"
	"git.wndv.co/sharp/app/models"
	"git.wndv.co/sharp/app/repos"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Init(rp *repos.Repos) *gin.Gin {
	r := gin.Default()
	r.Configurer().ProductionReady()
	e := repos.Init()
	r.GETw("/ping", func(ctx *gin.Context) (interface{}, error) {
		return gin.H{
			"message": "pong",
		}, nil
	})

	r.POSTw("/emails/register", func(ctx *gin.Context) (interface{}, error) {
		var email models.Email
		err := ctx.ShouldBindJSON(&email)
		if err != nil {
			return nil, err
		}
		email.ID = primitive.NewObjectID()
		err = e.Email.Insert(ctx, email)
		if err != nil {
			return nil, err
		}

		return gin.H{
			"result": email,
		}, nil
	})

	r.POSTw("/campaign/prepare", func(ctx *gin.Context) (interface{}, error) {
		var campaign models.Campaign
		err := ctx.ShouldBindJSON(&campaign)
		if err != nil {
			return nil, err
		}
		campaign.ID = primitive.NewObjectID()
		err = e.Campaign.Insert(ctx, campaign)
		if err != nil {
			return nil, err
		}

		return gin.H{
			"result": campaign.ID,
		}, nil
	})

	return r
}
