// Copyright 2016 The G3N Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package graphic

import (
	"github.com/kasworld/h4o/geometry"
	"github.com/kasworld/h4o/gls"
	"github.com/kasworld/h4o/material"
	"github.com/kasworld/h4o/renderinfo"
)

// Lines is a Graphic which is rendered as a collection of independent lines.
type Lines struct {
	Graphic             // Embedded graphic object
	uniMVPm gls.Uniform // Model view projection matrix uniform location cache
}

// NewLines returns a pointer to a new Lines object.
func NewLines(igeom geometry.GeometryI, imat material.MaterialI) *Lines {

	l := new(Lines)
	l.Init(igeom, imat)
	return l
}

// Init initializes the Lines object and adds the specified material.
func (l *Lines) Init(igeom geometry.GeometryI, imat material.MaterialI) {

	l.Graphic.Init(l, igeom, gls.LINES)
	l.AddMaterial(l, imat, 0, 0)
	l.uniMVPm.Init("MVP")
}

// RenderSetup is called by the engine before drawing this geometry.
func (l *Lines) RenderSetup(gs *gls.GLS, rinfo *renderinfo.RenderInfo) {

	// Transfer model view projection matrix uniform
	mvpm := l.ModelViewProjectionMatrix()
	location := l.uniMVPm.Location(gs)
	gs.UniformMatrix4fv(location, 1, false, &mvpm[0])
}
