package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/amscotti/urlRedis/storage"
	"github.com/go-zoo/bone"
)

// GetKey return a http.Handler that returns the key from storage
func GetKey(store storage.Database) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		key := bone.GetValue(r, "key")

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
