package physics

import (
	"time"

	"github.com/kasworld/h4o/_examples/app"
	"github.com/kasworld/h4o/appwindow"
	"github.com/kasworld/h4o/eventtype"
	"github.com/kasworld/h4o/experimental/physics"
	"github.com/kasworld/h4o/experimental/physics/object"
	"github.com/kasworld/h4o/geometry"
	"github.com/kasworld/h4o/graphic"
	"github.com/kasworld/h4o/light"
	"github.com/kasworld/h4o/material"
	"github.com/kasworld/h4o/math32"
	"github.com/kasworld/h4o/util/helper"
)

func init() {
	app.DemoMap["physics-experimental.sphere_box"] = &PhysicsSphereBox{}
}

type PhysicsSphereBox struct {
	sim *physics.Simulation
	rb  *object.Body
	rb2 *object.Body
	rb3 *object.Body
}

// Start is called once at the start of the demo.
func (t *PhysicsSphereBox) Start(a *app.App) {

	// Subscribe to key events
	a.Subscribe(eventtype.OnKeyRepeat, t.onKey)
	a.Subscribe(eventtype.OnKeyDown, t.onKey)

	// Create axes helper
	axes := helper.NewAxes(1)
	a.Scene().Add(axes)

	pl := light.NewPoint(math32.NewColor("white"), 1.0)
	pl.SetPosition(1, 0, 1)
	a.Scene().Add(pl)

	// Add directional green light from top
	l2 := light.NewDirectional(&math32.Color{1, 1, 1}, 0.3)
	l2.SetPosition(0, 0.1, 0)
	a.Scene().Add(l2)

	t.sim = physics.NewSimulation(a.Scene())

	sphereGeom := geometry.NewSphere(0.1, 16, 8)
	cubeGeom := geometry.NewCube(0.2)
	mat := material.NewStandard(&math32.Color{1, 1, 1})
	mat.SetTransparent(true)
	mat.SetOpacity(0.5)

	sphere := graphic.NewMesh(sphereGeom, mat)
	sphere.SetPosition(2, 0, 0)
	a.Scene().Add(sphere)
	t.rb2 = object.NewBody(sphere)
	//t.rb2.SetLinearDamping(0)
	t.sim.AddBody(t.rb2, "Sphere")
	t.rb2.SetVelocity(math32.NewVector3(-0.5, 0, 0))

	cube := graphic.NewMesh(cubeGeom, mat)
	cube.SetPosition(0, 0, 0)
	cube.SetRotation(0, math32.Pi*0.25, math32.Pi*0.25)
	a.Scene().Add(cube)
	t.rb3 = object.NewBody(cube)
	//t.rb3.SetLinearDamping(0)
	t.sim.AddBody(t.rb3, "Cube1")
	t.rb3.SetVelocity(math32.NewVector3(0.5, 0, 0))

}

func (t *PhysicsSphereBox) onKey(evname eventtype.EventType, ev interface{}) {

	kev := ev.(*appwindow.KeyEvent)
	switch kev.Key {
	case appwindow.KeyP:
		t.sim.SetPaused(!t.sim.Paused())
	case appwindow.KeySpace:
		t.sim.SetPaused(false)
		t.sim.Step(0.016)
		t.sim.SetPaused(true)
	case appwindow.Key1:
		t.rb2.ApplyVelocityDeltas(math32.NewVector3(-1, 0, 0), math32.NewVector3(0, 0, 1))
	case appwindow.Key2:
		t.rb2.ApplyVelocityDeltas(math32.NewVector3(1, 0, 0), math32.NewVector3(0, 0, -1))
	}
}

// Update is called every frame.
func (t *PhysicsSphereBox) Update(a *app.App, deltaTime time.Duration) {

	t.sim.Step(float32(deltaTime.Seconds()))
}

// Cleanup is called once at the end of the demo.
func (t *PhysicsSphereBox) Cleanup(a *app.App) {}
