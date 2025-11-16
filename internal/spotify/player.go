package spotify

import (
	"fmt"
	"github.com/godbus/dbus/v5"
	"os"
)

func OpenUri(uri string) {
	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to connect to session bus:", err)
		os.Exit(1)
	}
	defer conn.Close()

	bus := conn.Object("org.mpris.MediaPlayer2.spotify",
		"/org/mpris/MediaPlayer2")

	call := bus.Call("org.mpris.MediaPlayer2.Player.OpenUri", 0, uri)
	if call.Err != nil {
		panic(call.Err)
	}
}

func PlayPause() {
	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to connect to session bus:", err)
		os.Exit(1)
	}
	defer conn.Close()

	bus := conn.Object("org.mpris.MediaPlayer2.spotify",
		"/org/mpris/MediaPlayer2")

	call := bus.Call("org.mpris.MediaPlayer2.Player.PlayPause", 0)
	if call.Err != nil {
		panic(call.Err)
	}
}
