package controllers

import (
	"github.com/revel/revel"
)

type Player struct {
	*revel.Controller
}

func (c Player) Index() revel.Result {
	return c.Render()
}
