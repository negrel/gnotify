package gnotify

import (
	"github.com/godbus/dbus"
)

const (
	prefix    = "org.freedesktop.Notifications"
	notifPath = "/org/freedesktop/Notifications"
)

// Notifier is responsible for sending notification.
type Notifier struct {
	bo dbus.BusObject
}

// New return a notifier and connect it to the dbus.
func New() (*Notifier, error) {
	conn, err := dbus.SessionBus()

	return &Notifier{
		bo: conn.Object(prefix, notifPath),
	}, err
}

// NewFromConnection return a new notifier that use the given
// dbus connection.
func NewFromConnection(conn *dbus.Conn) *Notifier {
	return &Notifier{
		bo: conn.Object(prefix, notifPath),
	}
}

// GetServerInfo return the information on the server,
// the server name, vendor and version number.
func (n *Notifier) GetServerInfo() (name, vendor, version, specVersion string, err error) {
	call := n.bo.Call(
		"org.freedesktop.Notifications.GetServerInformation",
		0,
		name,
		vendor,
		version,
		specVersion,
	)
	err = call.Err

	return
}

// Notify send a notification to the notification server.
func (n *Notifier) Notify(notif Notification) {
	n.bo.Call(
		prefix+".Notify",
		0,
		notif.appName,
		notif.replacesID,
		notif.appIcon,
		notif.summary,
		notif.body,
		notif.actions,
		notif.hints,
		notif.expireTimeout,
	)

}
