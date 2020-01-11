package events

import (
	"fmt"
	"strings"

	"github.com/Equanox/gotron"
)

var v = make(chan gotron.Event)

func ExecChrome(window *gotron.BrowserWindow) {
	Evento := gotron.Event{Event: "Chrome"}

	window.On(&Evento, func(bin []byte) {
		variables := <-gotron.UserVar
		fmt.Println(strings.Split(variables.Variables, "|"))
	})

}
