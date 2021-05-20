package geometry

import (
	"time"

	"github.com/kasworld/h4o/_examples/app"
	"github.com/kasworld/h4o/geometry"
	"github.com/kasworld/h4o/gls"
	"github.com/kasworld/h4o/graphic"
	"github.com/kasworld/h4o/material"
	"github.com/kasworld/h4o/math32"
)

func init() {
	app.DemoMap["geometry.lines"] = &Lines{}
}

type Lines struct{}

// Start is called once at the start of the demo.
func (t *Lines) Start(a *app.App) {

	// Creates geometry
	geom := geometry.NewGeometry()
	vertices := math32.NewArrayF32(0, 16)
	vertices.Append(
		-0.5, 0.0, 0.0,
		0.5, 0.0, 0.0,
		0.0, -0.5, 0.0,
		0.0, 0.5, 0.0,
		0.0, 0.0, -0.5,
		0.0, 0.0, 0.5,
	)
	colors := math32.NewArrayF32(0, 16)
	colors.Append(
		1.0, 0.0, 0.0,
		1.0, 0.0, 0.0,
		0.0, 1.0, 0.0,
		0.0, 1.0, 0.0,
		0.0, 0.0, 1.0,
		0.0, 0.0, 1.0,
	)
	geom.AddVBO(gls.NewVBO(vertices).AddAttrib(gls.VertexPosition))
	geom.AddVBO(gls.NewVBO(colors).AddAttrib(gls.VertexColor))

	// Creates basic material
	mat := material.NewBasic()

	// Creates lines with the specified geometry and material
	lines1 := graphic.NewLines(geom, mat)
	a.Scene().Add(lines1)
}

// Update is called every frame.
func (t *Lines) Update(a *app.App, deltaTime time.Duration) {}

// Cleanup is called once at the end of the demo.
func (t *Lines) Cleanup(a *app.App) {}
