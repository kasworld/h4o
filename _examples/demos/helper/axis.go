package helper

import (
	"time"

	"github.com/kasworld/h4o/_examples/app"
	"github.com/kasworld/h4o/util/helper"
)

func init() {
	app.DemoMap["helper.axis"] = &AxisHelper{}
}

type AxisHelper struct{}

// Start is called once at the start of the demo.
func (t *AxisHelper) Start(a *app.App) {

	ah := helper.NewAxes(1.0)
	a.Scene().Add(ah)
}

// Update is called every frame.
func (t *AxisHelper) Update(a *app.App, deltaTime time.Duration) {}

// Cleanup is called once at the end of the demo.
func (t *AxisHelper) Cleanup(a *app.App) {}
