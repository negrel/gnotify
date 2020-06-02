package gnotify

import "time"

// // Event provide a generic interface for all events.
// type Event interface {
// 	// When the event was triggered.
// 	When() time.Time
// }

// var _ Event = CloseEvent{}

// Reason the notification was closed.
type Reason uint32

const (
	// ReasonExpired is used when the notification expired
	ReasonExpired Reason = iota + 1

	// ReasonDismissed is used when the notification dismissed
	ReasonDismissed

	// ReasonClosed is used when the notification is closed by the user
	ReasonClosed

	// ReasonUndefined is used for other reasons.
	ReasonUndefined
)

// var _ Event = CloseEvent{}

// CloseEvent are triggered when the user close the desktop
// notification.
type CloseEvent struct {
	When   time.Time
	Reason Reason
}

// When implements the Event interface.
//
// When return when the event was triggered.
// func (ce CloseEvent) When() time.Time {
// 	return ce.when
// }

// // Reason return the reason the notification is closed.
// func (ce CloseEvent) Reason() Reason {
// 	return ce.reason
// }

// CloseHandler handle notification close event
type CloseHandler func(CloseEvent)

// var _ Event = ActionEvent{}

// ActionEvent are triggered when the user click on one
// of the actions of the desktop notification.
type ActionEvent struct {
	When time.Time
	Key  string
}

// When implements the Event interface.
//
// When return when the event was triggered.
// func (ae ActionEvent) When() time.Time {
// 	return ae.when
// }

// ActionHandler handle notification action event
type ActionHandler func(ActionEvent)
