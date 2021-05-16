package animation

import "github.com/kasworld/h4o/math32"

// ChannelI is the interface for all channel types.
type ChannelI interface {
	Update(time float32)
	SetBuffers(keyframes, values math32.ArrayF32)
	Keyframes() math32.ArrayF32
	Values() math32.ArrayF32
	SetInterpolationType(it InterpolationType)
}
