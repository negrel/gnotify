package gnotify

// Tag is a notification identifier.
// Notification that share the same tag
// can be replaced by a new one if the
// Renotify option is set to true.
type Tag string

// NotificationOption is a custom object that contain setting that
// you want to apply to the notification.
type NotificationOption struct {
	title string

	// The URL of the image used to represent the notification
	// when there isn't enough space to display the
	// notification itself.
	Badge string

	// The body is the text of the notification, which
	// is displayed below the title.
	Body string

	// An identifying tag for the notification
	Tag Tag

	// The URL of the image used as an icon of the notification.
	Icon string

	// The URL of an image to be displayed as part of the notification.
	Image string

	// Renotify specify whether this notification should
	// replace the previous one.
	Renotify bool

	// The timeout time in milliseconds since the display of the notification
	// at which the notification should automatically close.
	ExpireTimeout int32

	// An array of NotificationActions representing the actions
	// available to the user when the notification is presented.
	// These are options the user can choose among in order
	// to act on the action within the context of the notification
	// itself.
	Actions []NotificationAction

	// A Boolean specifying whether the notification is silent.
	Silent bool
}
