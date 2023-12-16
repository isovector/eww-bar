package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/godbus/dbus/v5"
)

func main() {
	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = conn.Close()
		fmt.Println("Exiting....")
	}()

	call := conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"eavesdrop='true',interface='org.freedesktop.Notifications',member='Notify'")
	if call.Err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Failed to add match:", call.Err)
		os.Exit(1)
	}

	c := make(chan *dbus.Message, 10)
	conn.Eavesdrop(c)
	for v := range c {
		notification := decode(v.Body)
    if notification.AppIcon == "Spotify" {
      exec.Command("bash", "/home/sandy/.config/eww/scripts/music-updates").Start()
      fmt.Printf("%s\n", "song update")
		}
  }

}
