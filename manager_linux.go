// +build !windows

package gnotify

import (
	"fmt"
	"log"
	"strings"

	"github.com/godbus/dbus/v5"
)

const (
	prefix    = "org.freedesktop.Notifications"
	notifPath = "/org/freedesktop/Notifications"
)

type capability int8

const (
	actionIconsCapable capability = iota
	actionsCapable
	bodyCapable
	bodyHyperlinksCapable
	bodyImagesCapable
	bodyMarkupCapable
	iconMultiCapable
	iconStaticCapable
	persistenceCapable
	soundCapable
)

var capabilitiesMap = map[string]capability{
	"actions-icons":   actionIconsCapable,
	"actions":         actionsCapable,
	"body":            bodyCapable,
	"body-hyperlinks": bodyHyperlinksCapable,
	"body-images":     bodyImagesCapable,
	"body-markup":     bodyMarkupCapable,
	"icon-multi":      iconMultiCapable,
	"icon-static":     iconStaticCapable,
	"persistence":     persistenceCapable,
	"sound":           soundCapable,
}

var capabilities map[capability]bool = func() map[capability]bool {
	conn, _ := dbus.SessionBus()
	busObj := conn.Object(prefix, notifPath)

	// getting notification capabilities.
	call := busObj.Call(
		prefix+".GetCapabilities",
		0,
	)
	caps := strings.Split(fmt.Sprint(call.Body[0]), " ")
	caps = caps[1 : len(caps)-1]

	result := make(map[capability]bool, len(caps))

	for _, capability := range caps {
		result[capabilitiesMap[capability]] = true
	}

	return result
}()

type activeNotification struct {
	id           uint32
	tag          Tag
	closeHandler CloseHandler
	clickHandler ActionHandler
}

var _ Manager = &UnixManager{}

// UnixManager is the notification manager for UNIX systems.
type UnixManager struct {
	appName     string
	busObj      dbus.BusObject
	activeNotif []activeNotification
}

// New return a UnixManager connected to the dbus.
func New(appName string) (Manager, error) {
	conn, err := dbus.SessionBus()
	if err != nil {
		return nil, err
	}

	dbusSignal := make(chan *dbus.Signal)

	go func() {
		for {
			log.Println("Starting DBUS SIGNAL LISTENING")
			signal := <-dbusSignal

			log.Println(signal)

			switch signal.Name {
			case prefix + ".NotificationClosed":
				log.Printf("Notification closed %+v", signal)
			case prefix + ".ActionInvoked":
				log.Println("Action invoked", signal)
			default:
				log.Println(signal)
			}
		}
	}()

	// Listen to signals.
	conn.Signal(dbusSignal)

	return &UnixManager{
		appName:     appName,
		busObj:      conn.Object(prefix, notifPath),
		activeNotif: make([]activeNotification, 0),
	}, err
}

// Push the given notification to display it.
func (m *UnixManager) Push(notif *Notification) {
	var replaceID uint32
	var actions []string

	// Add actions if supported
	if capabilities[actionsCapable] {
		for _, action := range notif.Actions {
			actions = append(actions, action.Name, action.Name)
		}
	}

	// Replace previous notification with same tag
	if notif.Renotify {
		for i := len(m.activeNotif) - 1; i >= 0; i-- {
			activeNotif := m.activeNotif[i]

			// Tag are equal
			if activeNotif.tag == notif.Tag {
				replaceID = activeNotif.id
			}
		}
	}

	call := m.busObj.Call(
		prefix+".Notify",
		0,
		m.appName,
		replaceID,
		notif.Icon,
		notif.title,
		notif.Body,
		actions,
		map[string]interface{}{},
		notif.ExpireTimeout,
	)

	if len(call.Body) > 0 {
		id, ok := call.Body[0].(uint32)
		if !ok {
			return
		}

		// we keep only what we need.
		m.activeNotif = append(m.activeNotif, activeNotification{
			id:           id,
			tag:          notif.Tag,
			closeHandler: notif.OnClose,
			clickHandler: notif.OnClick,
		})
	}
}
