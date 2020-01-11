package hour

import (
	"encoding/json"
	"net/http"
	"time"
)

type Hora struct {
	Hora    int
	Minuto  int
	Segundo int
}

func RestAPI() {
	http.HandleFunc("/PrintHour", printHour)

	http.ListenAndServe("", nil)
}

func printHour(w http.ResponseWriter, r *http.Request) {
	h, m, s := time.Now().Clock()

	hour := Hora{Hora: h, Minuto: m, Segundo: s}

	Json, err := json.Marshal(hour)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(Json)
}
