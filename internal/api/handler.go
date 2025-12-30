package api

import (
	"encoding/json"
	"net/http"

	"go-rate-limiter/internal/limiter"
)

type Response struct {
	Allowed   bool `json:"allowed"`
	Remaining int  `json:"remaining"`
}

func AllowHandler(tb *limiter.TokenBucket) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		if key == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		allowed, remaining, _ := tb.Allow(r.Context(), key)

		json.NewEncoder(w).Encode(Response{
			Allowed:   allowed,
			Remaining: remaining,
		})
	}
}
