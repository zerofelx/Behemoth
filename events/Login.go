package events

import (
	"fmt"
	"strings"

	"github.com/zerofelx/Behemoth/database"

	"github.com/zerofelx/Behemoth/SendEvents"
	"github.com/zerofelx/Behemoth/views"

	"github.com/Equanox/gotron"
)

// User Logged
var ul = make(chan gotron.Event)
var verify = make(chan bool)

func LoginEvent(window *gotron.BrowserWindow) {
	Evento := gotron.Event{Event: "Login"}

	window.On(&Evento, func(bin []byte) {
		UserVariables := <-gotron.UserVar
		User := strings.Split(UserVariables.Variables, "|")
		go func() {
			verify <- database.QueryPassword(User[0], User[1])
		}()

		log := <-verify

		fmt.Println(log)
		if log == true {
			SendEvents.Logged(window)
			views.User = database.GetData(User[0])
		} else {
			SendEvents.LoginFailed(window)
		}
	})
}
