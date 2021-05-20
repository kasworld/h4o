package usrevent

import "github.com/kasworld/h4o/eventtype"

const (
	OnOK = eventtype.EventType(eventtype.EventType_Count + iota)
	OnCancel
	OnSelect
)
