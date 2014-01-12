package controllers

import (
	"github.com/robfig/revel"
	"github.com/mati1979/go/revel/revelhello/app/adstream"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	adstream.Init()
	return c.Render("")
}
