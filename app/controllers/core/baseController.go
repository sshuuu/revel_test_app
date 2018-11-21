package core

import (
	"github.com/revel/revel"
)

type BaseController struct {
	*revel.Controller
	Account struct {
		// models.Account
		// IsLogin bool
	}
}
