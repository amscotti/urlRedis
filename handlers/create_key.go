package handlers

import (
	"encoding/json"
	"html"
	"net/http"

	"github.com/amscotti/urlRedis/storage"
)

// CreateKey return a http.Handler that write a new key to storage
func CreateKey(store storage.Database) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		status, err := store.Set(html.EscapeString(r.FormValue("url")))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json, err := json.Marshal(status)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	}
}
