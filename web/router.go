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
		errS := ctx.ShouldBindJSON(&email)
		if errS != nil {
			return nil, errS
		}
		email.ID = primitive.NewObjectID()
		errI := e.Email.Insert(ctx, email)
		if errI != nil {
			return nil, errI
		}

		return gin.H{
			"result": email,
		}, nil
	})

	r.POSTw("/campaign/prepare", func(ctx *gin.Context) (interface{}, error) {
		var campaign models.Campaign
		errS := ctx.ShouldBindJSON(&campaign)
		if errS != nil {
			return nil, errS
		}
		campaign.ID = primitive.NewObjectID()
		errI := e.Campaign.Insert(ctx, campaign)
		if errI != nil {
			return nil, errI
		}

		return gin.H{
			"result": campaign.ID,
		}, nil
	})

	return r
}
