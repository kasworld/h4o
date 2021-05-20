package physics

import (
	"time"

	"github.com/kasworld/h4o/_examples/app"
	"github.com/kasworld/h4o/appwindow"
	"github.com/kasworld/h4o/eventtype"
	"github.com/kasworld/h4o/experimental/collision/shape"
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
	app.DemoMap["physics-experimental.basic"] = &PhysicsBasic{}
}

type PhysicsBasic struct {
	sim *physics.Simulation
	rb  *object.Body
	rb2 *object.Body
	rb3 *object.Body
}

// Start is called once at the start of the demo.
func (t *PhysicsBasic) Start(a *app.App) {

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
	//gravity := physics.NewConstantForceField(&math32.Vector3{0,-9.8,0})
	// //gravity := physics.NewAttractorForceField(&math32.Vector3{0.1,1,0}, 1)
	//t.sim.AddForceField(gravity)

	// Creates sphere 1
	sphereGeom := geometry.NewSphere(0.1, 16, 8)
	mat := material.NewStandard(&math32.Color{1, 1, 1})
	mat.SetWireframe(true)

	sphere1 := graphic.NewMesh(sphereGeom, mat)
	a.Scene().Add(sphere1)
	t.rb = object.NewBody(sphere1)
	t.rb.SetShape(shape.NewSphere(0.1)) // COMMENT THIS ON BOTH SPHERES TO TEST Convex-Convex collision!
	t.sim.AddBody(t.rb, "Sphere1")

	sphere2 := graphic.NewMesh(sphereGeom, mat)
	sphere2.SetPosition(1, 0, 0)
	a.Scene().Add(sphere2)
	t.rb2 = object.NewBody(sphere2)
	t.rb2.SetShape(shape.NewSphere(0.1)) // COMMENT THIS ON BOTH SPHERES TO TEST Convex-Convex collision!
	t.sim.AddBody(t.rb2, "Sphere2")
	t.rb2.SetVelocity(math32.NewVector3(-0.5, 0, 0))
	t.rb2.SetAngularVelocity(math32.NewVector3(0, 0, 1))

	//cubeGeom := geometry.NewCube(0.2)
	//cube1 := graphic.NewMesh(cubeGeom, mat)
	//a.Scene().Add(cube1)
	//t.rb3 = object.NewBody(cube1)
	//t.sim.AddBody(t.rb3, "Cube1")

}

func (t *PhysicsBasic) onKey(evname eventtype.EventType, ev interface{}) {

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
func (t *PhysicsBasic) Update(a *app.App, deltaTime time.Duration) {

	t.sim.Step(float32(deltaTime.Seconds()))
}

// Cleanup is called once at the end of the demo.
func (t *PhysicsBasic) Cleanup(a *app.App) {}
