package controllers

import "github.com/revel/revel"
import "github.com/matiwinnetou/go-revel-mobile-cars-adstream/app/adstream"

func init() {
	revel.OnAppStart(func() {
		adstream.Init()
	})
}
