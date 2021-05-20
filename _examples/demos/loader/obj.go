package loader

import (
	"path/filepath"
	"time"

	"github.com/kasworld/h4o/_examples/app"
	"github.com/kasworld/h4o/_examples/util"
	"github.com/kasworld/h4o/eventtype"
	"github.com/kasworld/h4o/light"
	"github.com/kasworld/h4o/loader/obj"
	"github.com/kasworld/h4o/math32"
	"github.com/kasworld/h4o/node"
	"github.com/kasworld/h4o/util/helper"
)

func init() {
	app.DemoMap["loader.obj"] = &LoaderObj{}
}

type LoaderObj struct {
	prevLoaded node.NodeI
	selFile    *util.FileSelectButton
}

// Start is called once at the start of the demo.
func (t *LoaderObj) Start(a *app.App) {

	// Creates file selection button
	t.selFile = util.NewFileSelectButton(a.DirData()+"/obj", "Select File", 400, 300)
	t.selFile.SetPosition(10, 10)
	t.selFile.FS.SetFileFilters("*.obj")
	a.DemoPanel().Add(t.selFile)
	t.selFile.Subscribe(eventtype.OnSelect, func(evname eventtype.EventType, ev interface{}) {
		fpath := ev.(string)
		err := t.load(a, fpath)
		if err == nil {
			t.selFile.Label.SetText("File: " + filepath.Base(fpath))
			t.selFile.SetError("")
		} else {
			t.selFile.Label.SetText("Select File")
		}
	})

	// Adds white directional front light
	l1 := light.NewDirectional(&math32.Color{1, 1, 1}, 1.0)
	l1.SetPosition(0, 0, 10)
	a.Scene().Add(l1)

	// Adds white directional top light
	l2 := light.NewDirectional(&math32.Color{1, 1, 1}, 1.0)
	l2.SetPosition(0, 10, 0)
	a.Scene().Add(l2)

	// Adds white directional right light
	l3 := light.NewDirectional(&math32.Color{1, 1, 1}, 1.0)
	l3.SetPosition(10, 0, 0)
	a.Scene().Add(l3)

	// Create axes helper
	axes := helper.NewAxes(2)
	a.Scene().Add(axes)

	fpath := "obj/cubemultitex.obj"
	t.load(a, filepath.Join(a.DirData(), fpath))
	t.selFile.Label.SetText("File: " + filepath.Base(fpath))
}

func (t *LoaderObj) load(a *app.App, path string) error {

	// Remove previous model from the scene
	if t.prevLoaded != nil {
		a.Scene().Remove(t.prevLoaded)
		t.prevLoaded.Dispose()
		t.prevLoaded = nil
	}

	// Decodes obj file and associated mtl file
	dec, err := obj.Decode(path, "")
	if err != nil {
		t.selFile.SetError(err.Error())
		return err
	}

	// Creates a new node with all the objects in the decoded file and adds it to the scene
	group, err := dec.NewGroup()
	if err != nil {
		t.selFile.SetError(err.Error())
		return err
	}
	a.Scene().Add(group)
	t.prevLoaded = group
	return nil
}

// Update is called every frame.
func (t *LoaderObj) Update(a *app.App, deltaTime time.Duration) {}

// Cleanup is called once at the end of the demo.
func (t *LoaderObj) Cleanup(a *app.App) {}
