package web

import (
	"git.wndv.co/go/srv/gin"
	"git.wndv.co/sharp/app/models"
	"git.wndv.co/sharp/app/repos"
)

func Init(rp *repos.Repos) *gin.Gin {
	r := gin.Default()
	r.Configurer().ProductionReady()
	repo := repos.Init()
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

		result, err := repo.Email.Insert(ctx, email)
		if err != nil {
			return nil, err
		}

		return gin.H{
			"result": result,
		}, nil
	})

	r.POSTw("/campaign/prepare", func(ctx *gin.Context) (interface{}, error) {
		var c models.Campaign
		err := ctx.ShouldBindJSON(&c)
		if err != nil {
			return nil, err
		}

		result, err := repo.Campaign.Insert(ctx, c)
		if err != nil {
			return nil, err
		}

		return gin.H{
			"result": result,
		}, nil
	})

	// r.POSTw("campaign/launch", func(ctx *gin.Context) (interface{}, error) {
	// 	var c models.Campaign
	// 	err := ctx.ShouldBindJSON(&c)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	emails, err := repo.Email.GetAll(ctx)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	for _, email := range emails {
	// 		str := services.FormatEmail(email, nil)
	// 	}

	// 	return gin.H{
	// 		"result": "",
	// 	}, nil
	// })

	return r
}
