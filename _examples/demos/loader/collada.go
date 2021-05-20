package loader

import (
	"io"
	"path/filepath"
	"time"

	"github.com/kasworld/h4o/_examples/app"
	"github.com/kasworld/h4o/_examples/util"
	"github.com/kasworld/h4o/eventtype"
	"github.com/kasworld/h4o/light"
	"github.com/kasworld/h4o/loader/collada"
	"github.com/kasworld/h4o/math32"
	"github.com/kasworld/h4o/node"
	"github.com/kasworld/h4o/util/helper"
)

func init() {
	app.DemoMap["loader.collada"] = &LoaderCollada{}
}

type LoaderCollada struct {
	prevLoaded  node.NodeI
	selFile     *util.FileSelectButton
	animTargets map[string]*collada.AnimationTarget
}

// Start is called once at the start of the demo.
func (t *LoaderCollada) Start(a *app.App) {

	// Creates file selection button
	t.selFile = util.NewFileSelectButton(a.DirData()+"/collada", "Select File", 400, 300)
	t.selFile.SetPosition(10, 10)
	t.selFile.FS.SetFileFilters("*.dae")
	a.DemoPanel().Add(t.selFile)
	t.selFile.Subscribe(eventtype.OnSelect, func(evname eventtype.EventType, ev interface{}) {
		fpath := ev.(string)
		err := t.load(a, fpath)
		if err == nil {
			t.selFile.Label.SetText("File: " + filepath.Base(fpath))
			t.selFile.SetError("")
		} else {
			t.selFile.Label.SetText("Select File")
			t.selFile.SetError(err.Error())
		}
	})

	// Add directional top white light
	l1 := light.NewDirectional(&math32.Color{1, 1, 1}, 1.0)
	l1.SetPosition(0, 1, 0)
	a.Scene().Add(l1)

	// Add directional right white light
	l2 := light.NewDirectional(&math32.Color{1, 1, 1}, 1.0)
	l2.SetPosition(1, 0, 0)
	a.Scene().Add(l2)

	// Add directional front  white light
	l3 := light.NewDirectional(&math32.Color{1, 1, 1}, 1.0)
	l3.SetPosition(0, 1, 1)
	a.Scene().Add(l3)

	// Sets camera position
	a.Camera().SetPosition(0, 4, 10)
	a.Camera().LookAt(&math32.Vector3{0, 0, 0}, &math32.Vector3{0, 1, 0})

	// Create axes helper
	axes := helper.NewAxes(1.5)
	a.Scene().Add(axes)

	// Loads default model
	fpath := filepath.Join(a.DirData(), "collada/scene.dae")
	t.load(a, fpath)
	t.selFile.Label.SetText("File: " + filepath.Base(fpath))
}

func (t *LoaderCollada) load(a *app.App, path string) error {

	// Remove previous model from the scene
	if t.prevLoaded != nil {
		a.Scene().Remove(t.prevLoaded)
		t.prevLoaded.Dispose()
		t.prevLoaded = nil
	}

	// Decodes collada file
	dec, err := collada.Decode(path)
	if err != nil && err != io.EOF {
		t.selFile.SetError(err.Error())
		return err
	}
	dec.SetDirImages(a.DirData() + "/images")

	// Loads collada scene
	s, err := dec.NewScene()
	if err != nil {
		t.selFile.SetError(err.Error())
		return err
	}
	a.Scene().Add(s)
	t.prevLoaded = s

	// Checks for animations
	ats, err := dec.NewAnimationTargets(s)
	if err == nil {
		t.animTargets = ats
		for _, at := range ats {
			at.SetStart(-1.0)
			at.Reset()
			at.SetLoop(true)
		}
	}
	return nil
}

// Update is called every frame.
func (t *LoaderCollada) Update(a *app.App, deltaTime time.Duration) {

	if t.animTargets != nil {
		dt := float32(deltaTime.Seconds())
		for _, at := range t.animTargets {
			at.Update(dt)
		}
	}
}

// Cleanup is called once at the end of the demo.
func (t *LoaderCollada) Cleanup(a *app.App) {}
