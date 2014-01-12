package controllers

import "github.com/robfig/revel"
import "github.com/mati1979/go-revel-mobile-cars-adstream/app/adstream"

func init() {
	revel.OnAppStart(func() {
		adstream.Init()
	})
}
