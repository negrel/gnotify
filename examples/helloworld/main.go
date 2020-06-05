package main

import (
	"log"

	gnotify "github.com/negrel/gnotify/pkg"
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
	notif := gnotify.Notification{
		Title: "Hello",
		Body:  "Gnotify is awesome.",
		Icon:  "/home/negrel/Documents/cni_front.jpg",
		Image: "/home/negrel/Documents/cni_back.jpg",
		// Expire after 3 seconds.
		ExpireTimeout: 3000,
	}

	manager.Push(&notif)
}
