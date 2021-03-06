// Copyright 2016 The G3N Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package orbit

import (
	"math"

	"github.com/kasworld/h4o/appwindow"
	"github.com/kasworld/h4o/camera"
	"github.com/kasworld/h4o/dispatcher"
	"github.com/kasworld/h4o/eventtype"
	"github.com/kasworld/h4o/gui"
	"github.com/kasworld/h4o/math32"
)

// OrbitEnabled specifies which control types are enabled.
type OrbitEnabled int

// The possible control types.
const (
	OrbitNone OrbitEnabled = 0x00
	OrbitRot  OrbitEnabled = 0x01
	OrbitZoom OrbitEnabled = 0x02
	OrbitPan  OrbitEnabled = 0x04
	OrbitKeys OrbitEnabled = 0x08
	OrbitAll  OrbitEnabled = 0xFF
)

// orbitState bitmask
type orbitState int

const (
	stateNone = orbitState(iota)
	stateRotate
	stateZoom
	statePan
)

// OrbitControl is a camera controller that allows orbiting a target point while looking at it.
// It allows the user to rotate, zoom, and pan a 3D scene using the mouse or keyboard.
type OrbitControl struct {
	dispatcher.Dispatcher                // Embedded event dispatcher
	cam                   *camera.Camera // Controlled camera
	target                math32.Vector3 // Camera target, around which the camera orbits
	up                    math32.Vector3 // The orbit axis (Y+)
	enabled               OrbitEnabled   // Which controls are enabled
	state                 orbitState     // Current control state

	// Public properties
	MinDistance     float32 // Minimum distance from target (default is 1)
	MaxDistance     float32 // Maximum distance from target (default is infinity)
	MinPolarAngle   float32 // Minimum polar angle in radians (default is 0)
	MaxPolarAngle   float32 // Maximum polar angle in radians (default is Pi)
	MinAzimuthAngle float32 // Minimum azimuthal angle in radians (default is negative infinity)
	MaxAzimuthAngle float32 // Maximum azimuthal angle in radians (default is infinity)
	RotSpeed        float32 // Rotation speed factor (default is 1)
	ZoomSpeed       float32 // Zoom speed factor (default is 0.1)
	KeyRotSpeed     float32 // Rotation delta in radians used on each rotation key event (default is the equivalent of 15 degrees)
	KeyZoomSpeed    float32 // Zoom delta used on each zoom key event (default is 2)
	KeyPanSpeed     float32 // Pan delta used on each pan key event (default is 35)

	// Internal
	rotStart  math32.Vector2
	panStart  math32.Vector2
	zoomStart float32
}

// NewOrbitControl creates and returns a pointer to a new orbit control for the specified camera.
func NewOrbitControl(cam *camera.Camera) *OrbitControl {

	oc := new(OrbitControl)
	oc.Dispatcher.Initialize()
	oc.cam = cam
	oc.target = *math32.NewVec3()
	oc.up = *math32.NewVector3(0, 1, 0)
	oc.enabled = OrbitAll

	oc.MinDistance = 1.0
	oc.MaxDistance = float32(math.Inf(1))
	oc.MinPolarAngle = 0
	oc.MaxPolarAngle = math32.Pi // 180 degrees as radians
	oc.MinAzimuthAngle = float32(math.Inf(-1))
	oc.MaxAzimuthAngle = float32(math.Inf(1))
	oc.RotSpeed = 1.0
	oc.ZoomSpeed = 0.1
	oc.KeyRotSpeed = 15 * math32.Pi / 180 // 15 degrees as radians
	oc.KeyZoomSpeed = 2.0
	oc.KeyPanSpeed = 35.0

	// Subscribe to events
	gui.Manager().SubscribeID(eventtype.OnMouseUp, &oc, oc.onMouse)
	gui.Manager().SubscribeID(eventtype.OnMouseDown, &oc, oc.onMouse)
	gui.Manager().SubscribeID(eventtype.OnScroll, &oc, oc.onScroll)
	gui.Manager().SubscribeID(eventtype.OnKeyDown, &oc, oc.onKey)
	gui.Manager().SubscribeID(eventtype.OnKeyRepeat, &oc, oc.onKey)
	oc.SubscribeID(eventtype.OnCursor, &oc, oc.onCursor)

	return oc
}

// Dispose unsubscribes from all events.
func (oc *OrbitControl) Dispose() {

	gui.Manager().UnsubscribeID(eventtype.OnMouseUp, &oc)
	gui.Manager().UnsubscribeID(eventtype.OnMouseDown, &oc)
	gui.Manager().UnsubscribeID(eventtype.OnScroll, &oc)
	gui.Manager().UnsubscribeID(eventtype.OnKeyDown, &oc)
	gui.Manager().UnsubscribeID(eventtype.OnKeyRepeat, &oc)
	oc.UnsubscribeID(eventtype.OnCursor, &oc)
}

// Reset resets the orbit control.
func (oc *OrbitControl) Reset() {

	oc.target = *math32.NewVec3()
}

// Target returns the current orbit target.
func (oc *OrbitControl) Target() math32.Vector3 {

	return oc.target
}

//Set camera orbit target Vector3
func (oc *OrbitControl) SetTarget(v math32.Vector3) {
	oc.target = v
}

// Enabled returns the current OrbitEnabled bitmask.
func (oc *OrbitControl) Enabled() OrbitEnabled {

	return oc.enabled
}

// SetEnabled sets the current OrbitEnabled bitmask.
func (oc *OrbitControl) SetEnabled(bitmask OrbitEnabled) {

	oc.enabled = bitmask
}

// Rotate rotates the camera around the target by the specified angles.
func (oc *OrbitControl) Rotate(thetaDelta, phiDelta float32) {

	const EPS = 0.0001

	// Compute direction vector from target to camera
	tcam := oc.cam.Position()
	tcam.Sub(&oc.target)

	// Calculate angles based on current camera position plus deltas
	radius := tcam.Length()
	theta := math32.Atan2(tcam.X, tcam.Z) + thetaDelta
	phi := math32.Acos(tcam.Y/radius) + phiDelta

	// Restrict phi and theta to be between desired limits
	phi = math32.Clamp(phi, oc.MinPolarAngle, oc.MaxPolarAngle)
	phi = math32.Clamp(phi, EPS, math32.Pi-EPS)
	theta = math32.Clamp(theta, oc.MinAzimuthAngle, oc.MaxAzimuthAngle)

	// Calculate new cartesian coordinates
	tcam.X = radius * math32.Sin(phi) * math32.Sin(theta)
	tcam.Y = radius * math32.Cos(phi)
	tcam.Z = radius * math32.Sin(phi) * math32.Cos(theta)

	// Update camera position and orientation
	oc.cam.SetPositionVec(oc.target.Clone().Add(&tcam))
	oc.cam.LookAt(&oc.target, &oc.up)
}

// Zoom moves the camera closer or farther from the target the specified amount
// and also updates the camera's orthographic size to match.
func (oc *OrbitControl) Zoom(delta float32) {

	// Compute direction vector from target to camera
	tcam := oc.cam.Position()
	tcam.Sub(&oc.target)

	// Calculate new distance from target and apply limits
	dist := tcam.Length() * (1 + delta/10)
	dist = math32.Max(oc.MinDistance, math32.Min(oc.MaxDistance, dist))
	tcam.SetLength(dist)

	// Update orthographic size and camera position with new distance
	oc.cam.UpdateSize(tcam.Length())
	oc.cam.SetPositionVec(oc.target.Clone().Add(&tcam))
}

// Pan pans the camera and target the specified amount on the plane perpendicular to the viewing direction.
func (oc *OrbitControl) Pan(deltaX, deltaY float32) {

	// Compute direction vector from camera to target
	position := oc.cam.Position()
	vdir := oc.target.Clone().Sub(&position)

	// Conversion constant between an on-screen cursor delta and its projection on the target plane
	c := 2 * vdir.Length() * math32.Tan((oc.cam.Fov()/2.0)*math32.Pi/180.0) / oc.winSize()

	// Calculate pan components, scale by the converted offsets and combine them
	var pan, panX, panY math32.Vector3
	panX.CrossVectors(&oc.up, vdir).Normalize()
	panY.CrossVectors(vdir, &panX).Normalize()
	panY.MultiplyScalar(c * deltaY)
	panX.MultiplyScalar(c * deltaX)
	pan.AddVectors(&panX, &panY)

	// Add pan offset to camera and target
	oc.cam.SetPositionVec(position.Add(&pan))
	oc.target.Add(&pan)
}

// onMouse is called when an OnMouseDown/OnMouseUp event is received.
func (oc *OrbitControl) onMouse(evname eventtype.EventType, ev interface{}) {

	// If nothing enabled ignore event
	if oc.enabled == OrbitNone {
		return
	}

	switch evname {
	case eventtype.OnMouseDown:
		gui.Manager().SetCursorFocus(oc)
		mev := ev.(*appwindow.MouseEvent)
		switch mev.Button {
		case appwindow.MouseButtonLeft: // Rotate
			if oc.enabled&OrbitRot != 0 {
				oc.state = stateRotate
				oc.rotStart.Set(mev.Xpos, mev.Ypos)
			}
		case appwindow.MouseButtonMiddle: // Zoom
			if oc.enabled&OrbitZoom != 0 {
				oc.state = stateZoom
				oc.zoomStart = mev.Ypos
			}
		case appwindow.MouseButtonRight: // Pan
			if oc.enabled&OrbitPan != 0 {
				oc.state = statePan
				oc.panStart.Set(mev.Xpos, mev.Ypos)
			}
		}
	case eventtype.OnMouseUp:
		gui.Manager().SetCursorFocus(nil)
		oc.state = stateNone
	}
}

// onCursor is called when an OnCursor event is received.
func (oc *OrbitControl) onCursor(evname eventtype.EventType, ev interface{}) {

	// If nothing enabled ignore event
	if oc.enabled == OrbitNone || oc.state == stateNone {
		return
	}

	mev := ev.(*appwindow.CursorEvent)
	switch oc.state {
	case stateRotate:
		c := -2 * math32.Pi * oc.RotSpeed / oc.winSize()
		oc.Rotate(c*(mev.Xpos-oc.rotStart.X),
			c*(mev.Ypos-oc.rotStart.Y))
		oc.rotStart.Set(mev.Xpos, mev.Ypos)
	case stateZoom:
		oc.Zoom(oc.ZoomSpeed * (mev.Ypos - oc.zoomStart))
		oc.zoomStart = mev.Ypos
	case statePan:
		oc.Pan(mev.Xpos-oc.panStart.X,
			mev.Ypos-oc.panStart.Y)
		oc.panStart.Set(mev.Xpos, mev.Ypos)
	}
}

// onScroll is called when an OnScroll event is received.
func (oc *OrbitControl) onScroll(evname eventtype.EventType, ev interface{}) {

	if oc.enabled&OrbitZoom != 0 {
		sev := ev.(*appwindow.ScrollEvent)
		oc.Zoom(-sev.Yoffset)
	}
}

// onKey is called when an OnKeyDown/OnKeyRepeat event is received.
func (oc *OrbitControl) onKey(evname eventtype.EventType, ev interface{}) {

	// If keyboard control is disabled ignore event
	if oc.enabled&OrbitKeys == 0 {
		return
	}

	kev := ev.(*appwindow.KeyEvent)
	if kev.Mods == 0 && oc.enabled&OrbitRot != 0 {
		switch kev.Key {
		case appwindow.KeyUp:
			oc.Rotate(0, -oc.KeyRotSpeed)
		case appwindow.KeyDown:
			oc.Rotate(0, oc.KeyRotSpeed)
		case appwindow.KeyLeft:
			oc.Rotate(-oc.KeyRotSpeed, 0)
		case appwindow.KeyRight:
			oc.Rotate(oc.KeyRotSpeed, 0)
		}
	}
	if kev.Mods == appwindow.ModControl && oc.enabled&OrbitZoom != 0 {
		switch kev.Key {
		case appwindow.KeyUp:
			oc.Zoom(-oc.KeyZoomSpeed)
		case appwindow.KeyDown:
			oc.Zoom(oc.KeyZoomSpeed)
		}
	}
	if kev.Mods == appwindow.ModShift && oc.enabled&OrbitPan != 0 {
		switch kev.Key {
		case appwindow.KeyUp:
			oc.Pan(0, oc.KeyPanSpeed)
		case appwindow.KeyDown:
			oc.Pan(0, -oc.KeyPanSpeed)
		case appwindow.KeyLeft:
			oc.Pan(oc.KeyPanSpeed, 0)
		case appwindow.KeyRight:
			oc.Pan(-oc.KeyPanSpeed, 0)
		}
	}
}

// winSize returns the appwindow height or width based on the camera reference axis.
func (oc *OrbitControl) winSize() float32 {

	width, size := appwindow.Get().GetSize()
	if oc.cam.Axis() == camera.Horizontal {
		size = width
	}
	return float32(size)
}
