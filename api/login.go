package api

import (
	"encoding/json"
	"net/http"
)

var Logeado = make(chan bool)

type UserLogged struct {
	Logeado  bool
	Username string
}

var Usuario = new(UserLogged)

func RunLogin(u string) {
	http.HandleFunc("/Logged", Logged)
	go func() { Logeado <- true }()
	Usuario.Username = u
	http.ListenAndServe("", nil)
}

func Logged(w http.ResponseWriter, r *http.Request) {
	lc := <-Logeado
	Usuario.Logeado = lc
	Json, err := json.Marshal(Usuario)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(Json)
}
