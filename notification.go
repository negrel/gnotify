package gnotify

import (
	"errors"
)

// // The optional name of the application sending the notification.
// appName string

// // The id of the notification to replace by this one.
// // O if no notification was replaced.
// replacesID ID

// // The optional program icon of the calling application.
// appIcon string

// // The summary text briefly describing the notification.
// title string

// // The optional detailed body text. Can be empty.
// message string

// // Actions are sent over as a list of pairs. Each even
// // element in the list (starting at index 0) represents
// // the identifier for the action. Each odd element in the
// // list is the localized string that will be displayed
// // to the user.
// actions []NotificationAction

// // Optional hints that can be passed to the server from
// // the client program.
// hints map[string]dbus.Variant

// // The timeout time in milliseconds since the display of
// // the notification at which the notification should
// // automatically close. If -1, the notification's expiration
// // time is dependent on the notification server's settings,
// // and may vary for the type of notification. If 0, never expire.
// expireTimeout int32

// // close handlers are responsible to handle close events.
// closeHandlers []func(CloseEvent)

// Notification is used to display desktop notifications
// to the user. Notification object must be passed to the manager.
type Notification struct {
	// settings of the notification.
	NotificationOption

	// Handle action events
	OnClick ActionHandler
	// Handle close events
	OnClose CloseHandler
}

// NewNotification return a reusable Notification object.
func NewNotification(title string, option NotificationOption) (*Notification, error) {
	if option.Renotify && option.Tag == "" {
		return nil, errors.New("notifications which set the renotify to true must specify a tag")
	}

	option.title = title

	return &Notification{
		NotificationOption: option,
	}, nil

	// return &Notification{
	// 	appName:       appName,
	// 	replacesID:    0,
	// 	appIcon:       "",
	// 	summary:       "",
	// 	message:       "",
	// 	hints:         make(map[string]dbus.Variant),
	// 	expireTimeout: 3000,
	// }, nil

}
