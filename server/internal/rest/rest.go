package rest

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/markovidakovic/surl/server/internal/config"
	"github.com/markovidakovic/surl/server/internal/db"
	"github.com/markovidakovic/surl/server/internal/rest/url"
	"gorm.io/gorm"
)

type Server struct {
	rtr *chi.Mux
	db  *gorm.DB
	cfg config.Config
}

func NewServer() *Server {
	cfg := config.Load()
	db := db.Connect(cfg)

	rtr := chi.NewRouter()
	rtr.Use(middleware.Logger)

	// Url resource
	urlRepo := url.NewRepository(db)
	urlService := url.NewService(urlRepo)
	urlHandler := url.NewHandler(urlService)

	rtr.Get("/urls", urlHandler.Get)

	return &Server{
		rtr,
		db,
		cfg,
	}
}

func (s *Server) Start() error {
	fmt.Printf("Server started on port %s\n", s.cfg.AppPort)
	return http.ListenAndServe(fmt.Sprintf(":%s", s.cfg.AppPort), s.rtr)
}
