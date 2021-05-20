package main

import (
	_ "github.com/kasworld/h4o/_examples/demos/animation"
	_ "github.com/kasworld/h4o/_examples/demos/audio"
	_ "github.com/kasworld/h4o/_examples/demos/experimental/physics"
	_ "github.com/kasworld/h4o/_examples/demos/geometry"
	_ "github.com/kasworld/h4o/_examples/demos/gui"
	_ "github.com/kasworld/h4o/_examples/demos/helper"
	_ "github.com/kasworld/h4o/_examples/demos/light"
	_ "github.com/kasworld/h4o/_examples/demos/loader"
	_ "github.com/kasworld/h4o/_examples/demos/material"
	_ "github.com/kasworld/h4o/_examples/demos/other"
	_ "github.com/kasworld/h4o/_examples/demos/shader"
	_ "github.com/kasworld/h4o/_examples/demos/tests"
	_ "github.com/kasworld/h4o/_examples/demos/texture"

	"github.com/kasworld/h4o/_examples/app"
)

func main() {

	// Create and run application
	app.Create().Run()
}
