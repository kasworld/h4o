package node

import (
	"github.com/kasworld/h4o/dispatcheri"
	"github.com/kasworld/h4o/gls"
	"github.com/kasworld/h4o/math32"
)

// NodeI is the interface for all node types.
type NodeI interface {
	dispatcheri.DispatcherI
	GetNode() *Node
	GetNodeI() NodeI
	Visible() bool
	SetVisible(state bool)
	Name() string
	SetName(string)
	Parent() NodeI
	Children() []NodeI
	IsAncestorOf(NodeI) bool
	LowestCommonAncestor(NodeI) NodeI
	UpdateMatrixWorld()
	BoundingBox() math32.Box3
	Render(gs *gls.GLS)
	Clone() NodeI
	Dispose()
	Position() math32.Vector3
	Rotation() math32.Vector3
	Scale() math32.Vector3
}
