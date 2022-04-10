package main

import (
	"git.wndv.co/go/pillars"
	"git.wndv.co/go/srv/grace"
	"git.wndv.co/sharp/app/repos"
	"git.wndv.co/sharp/app/web"
)

func main() {
	internalServer := pillars.NewInternalServer("email-campaign")
	go internalServer.RunDefault()

	rp := repos.Init()
	r := web.Init(rp)

	grace.OnSignal(func() {
		r.StopGracefully()
		internalServer.Stop()
	})

	r.Run("localhost:8080")
}
