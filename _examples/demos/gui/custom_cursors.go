package gui

import (
	"log"
	"time"

	"github.com/kasworld/h4o/_examples/app"
	"github.com/kasworld/h4o/appbase/appwindow"
	"github.com/kasworld/h4o/eventtype"
	"github.com/kasworld/h4o/gui"
)

func init() {
	app.DemoMap["gui.custom_cursors"] = &CustomCursors{}
}

type CustomCursors struct {
	cursors []appwindow.Cursor
	current int
}

// Start is called once at the start of the demo.
func (t *CustomCursors) Start(a *app.App) {

	// Show and enable demo panel
	a.DemoPanel().SetRenderable(true)
	a.DemoPanel().SetEnabled(true)

	t.cursors = make([]appwindow.Cursor, 2)
	t.current = 0

	instructions := gui.NewLabel("Click to change cursor!")
	instructions.SetPosition(50, 50)
	a.DemoPanel().Add(instructions)

	var err error
	t.cursors[0], err = a.CreateCursor(a.DirData()+"/images/gopher_cursor.png", 0, 0)
	t.cursors[1], err = a.CreateCursor(a.DirData()+"/images/gauntlet_cursor.png", 0, 0)
	if err != nil {
		log.Fatal("Error creating cursor: %s", err)
	}

	a.SetCursor(t.cursors[t.current])

	// Change cursor when clicking
	a.DemoPanel().SubscribeID(eventtype.OnMouseDown, a, func(s eventtype.EventType, i interface{}) {
		t.current += 1
		if t.current > len(t.cursors)-1 {
			t.current = 0
		}
		a.SetCursor(t.cursors[t.current])
	})

}

// Update is called every frame.
func (t *CustomCursors) Update(a *app.App, deltaTime time.Duration) {}

// Cleanup is called once at the end of the demo.
func (t *CustomCursors) Cleanup(a *app.App) {}
