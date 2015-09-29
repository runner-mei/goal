package controllers

import (
	h "net/http"

	"github.com/colegion/goal/commands/generate/handlers/testdata/controllers/subpackage"

	"github.com/naoina/denco"
	"github.com/revel/revel/testing"
)

// Controller is a struct that should be embedded into every controller
// of your app to make methods provided by middleware controllers available.
type Controller struct {
	*subpackage.Controller
	*denco.Param

	testing.TestSuite
	Test testing.TestSuite
}

// Before is a magic method that is executed before every request.
func (c *Controller) Before(uid string) h.Handler {
	return nil
}

// index is not an action as this method is not public.
func (c Controller) index(page int) h.Handler {
	return nil
}

// Index is a sample action.
func (c *App) Index(page int) h.Handler {
	return nil
}

// NotAction is not an action as this method doesn't return
// action.Result as its first argument.
func (c Controller) NotAction(page int) (bool, h.Handler) {
	return false, nil
}

// After is a magic method that is executed after every request.
func (c *Controller) After(name string) h.Handler {
	return nil
}