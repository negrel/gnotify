package gnotify

import "github.com/godbus/dbus"

// ID is a notification identifier.
type ID uint32

// Notification hold the notification information.
type Notification struct {
	// The optional name of the application sending the notification.
	appName string

	// The id of the notification to replace by this one.
	// O if no notification was replaced.
	replacesID ID

	// The optional program icon of the calling application.
	appIcon string

	// The summary text briefly describing the notification.
	summary string

	// The optional detailed body text. Can be empty.
	body string

	// Actions are sent over as a list of pairs. Each even
	// element in the list (starting at index 0) represents
	// the identifier for the action. Each odd element in the
	// list is the localized string that will be displayed
	// to the user.
	actions []string

	// Optional hints that can be passed to the server from
	// the client program.
	hints map[string]dbus.Variant

	// The timeout time in milliseconds since the display of
	// the notification at which the notification should
	// automatically close. If -1, the notification's expiration
	// time is dependent on the notification server's settings,
	// and may vary for the type of notification. If 0, never expire.
	expireTimeout int32
}

// NewNotification return a reusable Notification object.
func NewNotification(appName string) Notification {
	return Notification{
		appName:       appName,
		replacesID:    0,
		appIcon:       "",
		summary:       "",
		body:          "",
		actions:       make([]string, 0, 4),
		hints:         make(map[string]dbus.Variant),
		expireTimeout: 0,
	}
}

// Actions set the notification actions.
func (n Notification) Actions(actions ...string) Notification {
	n.actions = actions
	return n
}

// Body set the notification body.
func (n Notification) Body(b string) Notification {
	n.body = b
	return n
}

// Icon set the path to the notification icon.
func (n Notification) Icon(path string) Notification {
	n.appIcon = path
	return n
}

// Summary set the text that briefly describe the notification.
func (n Notification) Summary(s string) Notification {
	n.summary = s
	return n
}

// Title set the notification title.
// Alias for Summary.
func (n Notification) Title(t string) Notification {
	n.summary = t
	return n
}
