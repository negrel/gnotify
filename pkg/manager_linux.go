// +build linux

package gnotify

import (
	"time"

	"github.com/godbus/dbus/v5"
)

const (
	prefix    = "org.freedesktop.Notifications"
	notifPath = "/org/freedesktop/Notifications"
)

var _ Manager = &UnixManager{}

// UnixManager is the notification manager for UNIX systems.
type UnixManager struct {
	appName string
	busObj  dbus.BusObject
	actives map[uint32]*Notification
}

// New return a UnixManager connected to the dbus.
func New(appName string) (Manager, error) {
	conn, err := dbus.SessionBus()
	if err != nil {
		return nil, err
	}

	manager := &UnixManager{
		appName: appName,
		busObj:  conn.Object(prefix, notifPath),
		actives: make(map[uint32]*Notification, 8),
	}

	// Listen to notifications signals
	dbusSignal := make(chan *dbus.Signal)

	err = conn.AddMatchSignal(dbus.WithMatchSender(prefix))
	if err != nil {
		return nil, err
	}
	conn.Signal(dbusSignal)

	go manager.listenTo(dbusSignal)

	return manager, nil
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// ANCHOR Methods

func (m *UnixManager) displatchAction(signal *dbus.Signal) {
	id := signal.Body[0].(uint32)
	key := signal.Body[1].(string)

	m.actives[id].Actions.handle(key)
}

func (m *UnixManager) displatchClose(signal *dbus.Signal) {
	id := signal.Body[0].(uint32)
	reason := Reason(signal.Body[1].(uint32))

	for activeID, active := range m.actives {
		if activeID == id {
			delete(m.actives, id)

			closeEvent := CloseEvent{
				When:   time.Now(),
				Reason: reason,
			}
			active.OnClose(closeEvent)
		}
	}
}

// listen to dbus signal
func (m *UnixManager) listenTo(ch <-chan *dbus.Signal) {
	for {
		signal := <-ch

		switch signal.Name {
		case prefix + ".NotificationClosed":
			m.displatchClose(signal)

		case prefix + ".ActionInvoked":
			m.displatchAction(signal)

		default:
			continue
		}
	}
}

// Push the given notification to display it.
func (m *UnixManager) Push(notif *Notification) error {
	notif.applyDefault()

	call := m.busObj.Call(
		prefix+".Notify",
		0,
		m.appName,
		uint32(0),
		// notif.id, // replace id
		notif.Icon,
		notif.Title,
		notif.Body,
		notif.Actions.dbusFmt(),
		map[string]interface{}{},
		notif.ExpireTimeout,
	)

	// dbus error
	if call.Err != nil {
		return call.Err
	}

	// NOTE Replace id
	// if len(call.Body) > 0 {
	// 	id := call.Body[0].(uint32)

	// 	// we keep only what we need.
	// 	m.actives[id] = notif

	// 	return id, nil
	// }

	return nil
}
