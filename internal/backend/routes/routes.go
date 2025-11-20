package routes

import (
	"net/http"
)

type Router struct {
	Handler *handlers.PostHandler
}

func (r *Router) Setup() http.Handler {
	mux := http.NewServeMux()

	// Sessions
	mux.HandleFunc("POST /login", r.SessionHandler.LoginUser)
	mux.HandleFunc("POST /logout", r.SessionHandler.LogoutUser)
	mux.HandleFunc("POST /resend-verification", r.SessionHandler.ResendVerificationEmail)

	// Wrap with middleware
	return r.AuthMiddleware.Handler(mux)
}
