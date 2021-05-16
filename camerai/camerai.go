package camerai

import "github.com/kasworld/h4o/math32"

// CameraI is the interface for all cameras.
type CameraI interface {
	ViewMatrix(m *math32.Matrix4)
	ProjMatrix(m *math32.Matrix4)
}
