package gnotify

// NotificationAction define a notification action.
type NotificationAction struct {
	Name    string
	Handler ActionHandler
}
