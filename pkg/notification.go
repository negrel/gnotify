package gnotify

// Notification is used to display desktop notifications
// to the user. Notification are displayed by
// pushing them to the Manager.
type Notification struct {
	id uint32

	// settings of the notification.
	Option
}

// NewNotification return a Notification object.
func NewNotification(option Option) *Notification {
	if option.OnClose == nil {
		option.OnClose = func(_ CloseEvent) {}
	}

	if option.Actions == nil {
		option.Actions = make(ActionList, 8)
	}

	return &Notification{
		Option: option,
	}
}
