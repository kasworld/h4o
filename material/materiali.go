package material

import "github.com/kasworld/h4o/gls"

// MaterialI is the interface for all materials.
type MaterialI interface {
	GetMaterial() *Material
	RenderSetup(gs *gls.GLS)
	Dispose()
}
