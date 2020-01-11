package SendEvents

import (
	"fmt"

	"github.com/Equanox/gotron"
)

type SEvent struct {
	*gotron.Event
	CustomAttribute string `json:"AtrNameInFrontend"`
}

func SendEvent(window *gotron.BrowserWindow) {
	fmt.Println()
	fmt.Println("Enviando evento")
	fmt.Println()
	window.Send(&SEvent{
		Event:           &gotron.Event{Event: "Test", Variables: "Zero"},
		CustomAttribute: "Hello world",
	})
}
