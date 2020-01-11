package sources

import "net/http"

func ServeSources() {
	http.Handle("/sources/", http.StripPrefix("/sources/", http.FileServer(http.Dir("Assets"))))

	http.ListenAndServe("", nil)
}
