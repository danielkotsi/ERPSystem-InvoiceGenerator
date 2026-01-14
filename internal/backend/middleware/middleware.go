package middleware

import (
	"context"
	"net/http"
	"time"
)

type Middleware struct {
}

func NewMiddleware() *Middleware {
	return &Middleware{}
}

func (m *Middleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx, cancel := context.WithTimeout(r.Context(), 25*time.Second)
		defer cancel()

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
