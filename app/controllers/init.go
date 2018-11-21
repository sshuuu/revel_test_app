package controllers

import (
	"test_app/app/models"

	"github.com/revel/revel"
)

// application interceptor, it called only once when application begin
func init() {
	println("init")
	revel.OnAppStart(models.InitDB)
	revel.InterceptFunc(checkUser, revel.BEFORE, &App{})
}

func checkUser(c *revel.Controller) revel.Result {
	println("checkUser")
	return nil
}
