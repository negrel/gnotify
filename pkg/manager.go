package gnotify

// Manager is responsible to display desktop notification.
type Manager interface {
	Push(*Notification)
}
