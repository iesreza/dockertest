package helloworld

import (
	"github.com/iesreza/foundation/lib/router"
	"github.com/iesreza/foundation/system"
)

func (component component) Routers() {

	system.Router.Match("hello","*", func(req router.Request) {
		req.WriteString("hello world")
	})

	system.Router.Static(component.Assets,component.Assets,nil)
	system.Router.Fallback = func(req router.Request) {
		req.WriteObject(struct{
			Success bool
			Error string
		}{
			Success:false,
			Error:"404 path not found",
		})
	}
}
