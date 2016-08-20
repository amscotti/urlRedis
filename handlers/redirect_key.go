package handlers

import (
	"html"
	"net/http"

	"github.com/amscotti/urlRedis/storage"
	"github.com/go-zoo/bone"
)

// RedirectKey return a http.Handler that redirects to the URL in the store
func RedirectKey(store storage.Database) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		key := bone.GetValue(r, "key")

		status, err := store.Get(key)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			store.Incr(key)
			http.Redirect(w, r, html.EscapeString(status.URL), 301)
		}
	}
}
