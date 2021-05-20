package physics

import (
	"log"
	"time"

	"github.com/kasworld/h4o/_examples/app"
	"github.com/kasworld/h4o/appwindow"
	"github.com/kasworld/h4o/eventtype"
	"github.com/kasworld/h4o/experimental/physics"
	"github.com/kasworld/h4o/experimental/physics/object"
	"github.com/kasworld/h4o/geometry"
	"github.com/kasworld/h4o/gls"
	"github.com/kasworld/h4o/graphic"
	"github.com/kasworld/h4o/light"
	"github.com/kasworld/h4o/material"
	"github.com/kasworld/h4o/math32"
	"github.com/kasworld/h4o/texture"
	"github.com/kasworld/h4o/util/helper"
)

func init() {
	app.DemoMap["physics-experimental.spheres2"] = &PhysicsSpheres2{}
}

type PhysicsSpheres2 struct {
	app        *app.App
	sim        *physics.Simulation
	sphereGeom *geometry.Geometry
	matSphere  *material.Standard
}

// Start is called once at the start of the demo.
func (t *PhysicsSpheres2) Start(a *app.App) {

	t.app = a

	// Subscribe to key events
	a.Subscribe(eventtype.OnKeyRepeat, t.onKey)
	a.Subscribe(eventtype.OnKeyDown, t.onKey)

	// Create axes helper
	axes := helper.NewAxes(1)
	a.Scene().Add(axes)

	pl := light.NewPoint(math32.NewColor("white"), 1.0)
	pl.SetPosition(1, 0, 1)
	a.Scene().Add(pl)

	// Add directional light from top
	l2 := light.NewDirectional(&math32.Color{1, 1, 1}, 0.3)
	l2.SetPosition(0, 0.1, 0)
	a.Scene().Add(l2)

	// Add directional light from top
	l3 := light.NewDirectional(&math32.Color{1, 1, 1}, 0.3)
	l3.SetPosition(0.1, 0, 0.1)
	a.Scene().Add(l3)

	t.sim = physics.NewSimulation(a.Scene())
	gravity := physics.NewConstantForceField(&math32.Vector3{0, -0.98, 0})
	// //gravity := physics.NewAttractorForceField(&math32.Vector3{0.1,1,0}, 1)
	t.sim.AddForceField(gravity)

	// Creates sphere 1
	t.sphereGeom = geometry.NewSphere(0.1, 16, 8)

	texfileG := a.DirData() + "/images/ground2.jpg"
	texG, err := texture.NewTexture2DFromImage(texfileG)
	texG.SetRepeat(10, 10)
	texG.SetWrapS(gls.REPEAT)
	texG.SetWrapT(gls.REPEAT)
	if err != nil {
		log.Fatal("Error loading texture: %s", err)
	}

	mat := material.NewStandard(&math32.Color{1, 1, 1})
	mat.SetTransparent(true)
	mat.SetOpacity(0.5)
	mat.AddTexture(texG)
	//mat.SetWireframe(true)

	//sphere1 := graphic.NewMesh(sphereGeom, mat)
	//a.Scene().Add(sphere1)
	//t.rb = object.NewBody(sphere1)
	//t.sim.AddBody(t.rb, "Sphere1")

	floorGeom := geometry.NewBox(10, 0.5, 10)
	floor := graphic.NewMesh(floorGeom, mat)
	floor.SetPosition(3, -0.2, 0)
	a.Scene().Add(floor)
	floorBody := object.NewBody(floor)
	floorBody.SetBodyType(object.Static)
	t.sim.AddBody(floorBody, "Floor")

	// Creates texture 3
	texfile := a.DirData() + "/images/uvgrid.jpg"
	tex3, err := texture.NewTexture2DFromImage(texfile)
	if err != nil {
		log.Fatal("Error loading texture: %s", err)
	}
	//tex3.SetFlipY(false)
	// Creates sphere 3
	t.matSphere = material.NewStandard(&math32.Color{1, 1, 1})
	t.matSphere.AddTexture(tex3)

	sphere2 := graphic.NewMesh(t.sphereGeom, t.matSphere)
	sphere2.SetPosition(0, 1, -0.02)
	a.Scene().Add(sphere2)
	rb2 := object.NewBody(sphere2)
	t.sim.AddBody(rb2, "Sphere2")

	sphere3 := graphic.NewMesh(t.sphereGeom, t.matSphere)
	sphere3.SetPosition(0.05, 1.2, 0.05)
	a.Scene().Add(sphere3)
	rb3 := object.NewBody(sphere3)
	t.sim.AddBody(rb3, "Sphere3")

	sphere4 := graphic.NewMesh(t.sphereGeom, t.matSphere)
	sphere4.SetPosition(-0.05, 1.4, 0)
	a.Scene().Add(sphere4)
	rb4 := object.NewBody(sphere4)
	t.sim.AddBody(rb4, "Sphere4")
}

func (t *PhysicsSpheres2) ThrowBall() {

	camPos := t.app.Camera().Position()
	camTarget := t.app.Orbit().Target()
	throwDir := math32.NewVec3().SubVectors(&camTarget, &camPos).SetLength(3)

	sphere := graphic.NewMesh(t.sphereGeom, t.matSphere)
	sphere.SetPositionVec(&camPos)
	t.app.Scene().Add(sphere)
	rb := object.NewBody(sphere)
	rb.SetVelocity(throwDir)
	t.sim.AddBody(rb, "Sphere4")
}

func (t *PhysicsSpheres2) onKey(evname eventtype.EventType, ev interface{}) {

	kev := ev.(*appwindow.KeyEvent)
	switch kev.Key {
	case appwindow.KeyP:
		t.sim.SetPaused(!t.sim.Paused())
	case appwindow.KeyO:
		t.sim.SetPaused(false)
		t.sim.Step(0.016)
		t.sim.SetPaused(true)
	case appwindow.KeySpace:
		t.ThrowBall()
	case appwindow.Key1:
		// TODO
	case appwindow.Key2:
		// TODO
	}
}

// Update is called every frame.
func (t *PhysicsSpheres2) Update(a *app.App, deltaTime time.Duration) {

	t.sim.Step(float32(deltaTime.Seconds()))
}

// Cleanup is called once at the end of the demo.
func (t *PhysicsSpheres2) Cleanup(a *app.App) {}
