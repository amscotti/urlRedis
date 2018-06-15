package handlers

import (
	"encoding/json"
	"html"
	"net/http"

	"github.com/amscotti/urlRedis/storage"
	"github.com/gorilla/mux"
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

// GetKey return a http.Handler that returns the key from storage
func GetKey(store storage.Database) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		key := vars["key"]

		status, err := store.Get(key)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
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

// RedirectKey return a http.Handler that redirects to the URL in the store
func RedirectKey(store storage.Database) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		key := vars["key"]

		status, err := store.Get(key)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			store.Incr(key)
			http.Redirect(w, r, html.EscapeString(status.URL), 301)
		}
	}
}
