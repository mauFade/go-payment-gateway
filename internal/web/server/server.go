package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mauFade/go-payment-gateway/internal/service"
	"github.com/mauFade/go-payment-gateway/internal/web/handlers"
)

type Server struct {
	router         *chi.Mux
	server         *http.Server
	accountService *service.AccountService
	port           string
}

func NewServer(ac *service.AccountService, port string) *Server {
	return &Server{
		accountService: ac,
		port:           port,
		router:         chi.NewRouter(),
	}
}

func (s *Server) ConfigureRoutes() {
	accountHandler := handlers.NewAccountHandler(*s.accountService)

	s.router.Post("/accounts", accountHandler.Create)
	s.router.Get("/accounts", accountHandler.Get)
}

func (s *Server) Start() error {
	s.server = &http.Server{
		Addr:    ":" + s.port,
		Handler: s.router,
	}
	return s.server.ListenAndServe()
}
