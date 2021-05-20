package other

import (
	"time"

	"github.com/kasworld/h4o/_examples/app"
	"github.com/kasworld/h4o/geometry"
	"github.com/kasworld/h4o/graphic"
	"github.com/kasworld/h4o/log"
	"github.com/kasworld/h4o/math32"

	"math/rand"

	"github.com/kasworld/h4o/material"
)

func init() {
	app.DemoMap["other.performance"] = &Performance{}
}

type Performance struct{}

// Start is called once at the start of the demo.
func (t *Performance) Start(a *app.App) {

	torusGeometry := geometry.NewTorus(0.5, 0.2, 16, 16, 2*math32.Pi)

	halfSize := 20
	step := 2
	count := 0
	for i := -halfSize; i < (halfSize + 1); i += step {
		for j := -halfSize; j < (halfSize + 1); j += step {
			for k := -halfSize; k < (halfSize + 1); k += step {
				count += 1
				mat := material.NewStandard(&math32.Color{rand.Float32(), rand.Float32(), rand.Float32()})
				//mat.SetSpecularColor(math32.NewColor("white"))
				//mat.SetShininess(50)
				torus := graphic.NewMesh(torusGeometry, mat)
				torus.SetPosition(float32(i), float32(j), float32(k))
				torus.SetRotation(rand.Float32()*2*math32.Pi, rand.Float32()*2*math32.Pi, rand.Float32()*2*math32.Pi)
				//torus.Materials()[0].GetMaterial().GetMaterial().SetWireframe(true)
				a.Scene().Add(torus)
			}
		}
	}
	log.Info("%v objects added to the scene!", count)

	//stepLight := 10
	//countLight := 0
	//for i := -halfSize; i < (halfSize+1); i+=stepLight {
	//	for j := -halfSize; j < (halfSize+1); j+=stepLight {
	//		for k := -halfSize; k < (halfSize+1); k+=stepLight {
	//			countLight += 1
	//			light := light.NewPoint(math32.NewColor("white"), 2.0)
	//			light.SetPosition(float32(i), float32(j), float32(k))
	//			light.SetLinearDecay(0.5)
	//			light.SetQuadraticDecay(0.5)
	//			a.Scene().Add(light)
	//		}
	//	}
	//}
	//log.Info("%v lights added to the scene!", countLight)

}

// Update is called every frame.
func (t *Performance) Update(a *app.App, deltaTime time.Duration) {}

// Cleanup is called once at the end of the demo.
func (t *Performance) Cleanup(a *app.App) {}
