package graphic

import (
	"github.com/kasworld/h4o/geometry"
	"github.com/kasworld/h4o/gls"
	"github.com/kasworld/h4o/node"
	"github.com/kasworld/h4o/renderinfo"
)

// GraphicI is the interface for all Graphic objects.
type GraphicI interface {
	node.NodeI
	GetGraphic() *Graphic
	GetGeometry() *geometry.Geometry
	GeometryI() geometry.GeometryI
	SetRenderable(bool)
	Renderable() bool
	SetCullable(bool)
	Cullable() bool
	RenderSetup(gs *gls.GLS, rinfo *renderinfo.RenderInfo)
}
