package gnotify

// Tag is a notification identifier.
// Notification that share the same tag
// can be replaced by a new one if the
// Renotify option is set to true.
type Tag string

// ActionList is a list of notification
// action.
type ActionList map[string]func()

// format action list for dbus
func (al ActionList) dbusFmt() []string {
	result := make([]string, 0, len(al)*2)

	for name := range al {
		result = append(result, name, name)
	}

	return result
}

func (al ActionList) handle(key string) {
	al[key]()
}

// Notification is used to display desktop notifications
// to the user. Notification are displayed by
// pushing them to the Manager.
type Notification struct {

	// Title of the notification.
	Title string

	// The URL of the image used to represent the notification
	// when there isn't enough space to display the
	// notification itself.
	Badge string

	// The body is the text of the notification, which
	// is displayed below the title.
	Body string

	// The URL of the image used as an icon of the notification.
	Icon string

	// The URL of an image to be displayed as part of the notification.
	Image string

	// The timeout time in milliseconds since the display of the notification
	// at which the notification should automatically close.
	ExpireTimeout int32

	// An array of NotificationActions representing the actions
	// available to the user when the notification is presented.
	// These are options the user can choose among in order
	// to act on the action within the context of the notification
	// itself.
	Actions ActionList

	// OnClose is triggered when the notifcation is closed
	OnClose func(CloseEvent)

	// A Boolean specifying whether the notification is silent.
	Silent bool
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// ANCHOR Methods

func (n *Notification) applyDefault() {
	if n.OnClose == nil {
		n.OnClose = func(_ CloseEvent) {}
	}

	if n.Actions == nil {
		n.Actions = map[string]func(){}
	}
}
