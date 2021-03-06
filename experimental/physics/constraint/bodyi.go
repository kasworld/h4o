package constraint

import (
	"github.com/kasworld/h4o/experimental/physics/equation"
	"github.com/kasworld/h4o/math32"
)

type BodyI interface {
	equation.BodyI
	WakeUp()
	VectorToWorld(*math32.Vector3) math32.Vector3
	PointToLocal(*math32.Vector3) math32.Vector3
	VectorToLocal(*math32.Vector3) math32.Vector3
	Quaternion() *math32.Quaternion
}

type ConstraintI interface {
	Update() // Update all the equations with data.
	Equations() []equation.EquationI
	CollideConnected() bool
	BodyA() BodyI
	BodyB() BodyI
}
