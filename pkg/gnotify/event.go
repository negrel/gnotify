package gnotify

import "time"

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

// CloseEvent are triggered when the user close the desktop
// notification.
type CloseEvent struct {
	When   time.Time
	Reason Reason
}

// ActionEvent are triggered when the user click on one
// of the actions of the desktop notification.
type ActionEvent struct {
	When time.Time
	Key  string
}
