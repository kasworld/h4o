package physics

import "github.com/kasworld/h4o/math32"

// ForceFieldI represents a force field. A force is defined for every point.
type ForceFieldI interface {
	ForceAt(pos *math32.Vector3) math32.Vector3
}
