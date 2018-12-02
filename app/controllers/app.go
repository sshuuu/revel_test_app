package controllers

import (
	"github.com/revel/revel"
)

// App Structure
type App struct {
	GormController
}

// Before is interceptor called before method
func (c App) Before() (result revel.Result, controller App) {
	println("Before method in App Controller")
	return result, controller
}

// After is interceptor called after method
func (c App) After() (result revel.Result, controller App) {
	println("After method in App Controller")
	return result, controller
}

// Index show top page
func (c App) Index() revel.Result {
	greeting := "Aloha World"
	return c.Render(greeting)
	// return c.Render()
}

// Hello say hello
func (c App) Hello(myName string) revel.Result {
	c.Validation.Required(myName).Message("Your name is required!")
	c.Validation.MinSize(myName, 3).Message("Your name is not long enough!")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}
	return c.Render(myName)
}
