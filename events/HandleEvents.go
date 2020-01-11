package events

import (
	"fmt"

	"github.com/Equanox/gotron"
)

func HandleEvents(window *gotron.BrowserWindow) {
	Evento := gotron.Event{Event: "Log"}

	window.On(&Evento, func(bin []byte) {
		fmt.Println("Evento")
	})
}
