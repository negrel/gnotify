package main

import (
	"log"

	"github.com/negrel/gnotify"
)

var manager gnotify.Manager

func init() {
	var err error
	manager, err = gnotify.New("Example")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	notif := gnotify.NewNotification(gnotify.Option{
		Title: "Hello",
		Body:  "Gnotify is awesome.",
		// Expire after 3 seconds.
		ExpireTimeout: 3000,
	})

	manager.Push(notif)
}
