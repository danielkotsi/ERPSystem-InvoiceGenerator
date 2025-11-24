package middleware

import (
	"context"
	"-invoice_manager/internal/backend/models"
	"net/http"
	"time"
)

type Middleware struct {
	Config *models.Config
}

func NewMiddleware(conf *models.Config) *Middleware {
	return &Middleware{Config: conf}
}

func (m *Middleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx, cancel := context.WithTimeout(r.Context(), 25*time.Second)
		defer cancel()

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
