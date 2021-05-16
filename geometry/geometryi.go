package geometry

import "github.com/kasworld/h4o/gls"

// GeometryI is the interface for all geometries.
type GeometryI interface {
	GetGeometry() *Geometry
	RenderSetup(gs *gls.GLS)
	Dispose()
}
