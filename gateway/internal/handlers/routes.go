package handlers

import (
	"github.com/gorilla/mux"
	"github.com/mynreden/microservices-go/gateway/internal/middleware"
	"net/http"
)

func (h *Handler) Routes() http.Handler {
	r := mux.NewRouter()
	r.Use(middleware.Logger)
	r.PathPrefix("/api/users/").Handler(http.StripPrefix("/api/users", h.userHandler()))
	r.PathPrefix("/api/posts/").Handler(http.StripPrefix("/api/posts", h.postHandler()))
	return r
}
