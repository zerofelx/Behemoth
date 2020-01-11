package views

import (
	"fmt"
	"time"

	"github.com/zerofelx/Behemoth/database"
)

var User = new(database.Data)

func VerifyLogin() {
	time.Sleep(15 * time.Second)
	fmt.Println(User)
}
