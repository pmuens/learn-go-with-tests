package context

import (
	"context"
	"fmt"
	"net/http"
)

// Store fetches data.
type Store interface {
	Fetch(ctx context.Context) (string, error)
}

// Server returns a handler for calling Store.
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			return
		}

		fmt.Fprint(w, data)
	}
}
