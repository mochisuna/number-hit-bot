package handler

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/mochisuna/number-hit-bot/domain/service"
	"github.com/unrolled/render"
)

// Server HTTP server
type Server struct {
	server *http.Server
	controller *adapter.Controller
	// *Services
	// *LineBot
}

// New server
func New(addr string, ctrs *adapter.Controller) *Server {
	return &Server{
		server: &http.Server{
			Addr: addr,
		},
		controller: ctrs,
		// Services: services,
		// LineBot:  line,
	}
}

// ListenAndServe override http ListenAndServe
func (s *Server) ListenAndServe() error {
	r := chi.NewRouter()

	// cord option
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
	})
	r.Use(cors.Handler)

	// 公式提供のmiddleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	// routings
	// 開発レベルでバージョン分けする可能性がゼロではないので一応バージョンをラベル切っておく
	r.Route("/v1", func(r chi.Router) {
		r.Post("/callback", s.controller.Callback)
	})
	// health check
	r.Route("/health", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
	})

	s.server.Handler = r
	return s.server.ListenAndServe()
}
