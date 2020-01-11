package SendEvents

import (
	"github.com/Equanox/gotron"
)

type UserLogged struct {
	*gotron.Event
	Logged bool `json:"logged"`
}

func Logged(window *gotron.BrowserWindow) {
	window.Send(&UserLogged{
		Event:  &gotron.Event{Event: "Logged"},
		Logged: true,
	})
}

func LoginFailed(window *gotron.BrowserWindow) {
	window.Send(&UserLogged{
		Event:  &gotron.Event{Event: "Logged"},
		Logged: false,
	})
}
