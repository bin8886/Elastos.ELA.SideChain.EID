package dpos_msg

import "github.com/elastos/Elastos.ELA/events"

// Constants for the type of a notification message.
const (
	ETOnProposal events.EventType = 2001
	ETSelfOnDuty events.EventType = 2002
	ETOnArbiter events.EventType = 2003
)