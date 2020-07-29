// +build windows

package gnotify

import (
	"github.com/go-toast/toast"
)

var _ Manager = &WindowsManager{}

// WindowsManager is the notification manager for Windows system.
type WindowsManager struct {
	appName string
}

// New return a new Notification manager.
func New(appName string) (Manager, error) {
	return &WindowsManager{
		appName: appName,
	}, nil
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// ANCHOR Methods

// Push the given notification to display it.
func (m *WindowsManager) Push(notif *Notification) error {
	actions := func() []toast.Action {
		list := make([]toast.Action, 0, len(notif.Actions))

		for name := range notif.Actions {
			list = append(list, toast.Action{
				Type:      "protocol",
				Label:     name,
				Arguments: "",
			})
		}

		return list
	}()

	toast := toast.Notification{
		AppID:   m.appName,
		Title:   notif.Title,
		Message: notif.Body,
		Icon:    notif.Icon,
		Actions: actions,
	}

	err := toast.Push()

	return err
}
