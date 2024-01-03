package apiserver

import (
	"net/http"

	"github.com/AlexCorn999/order-data-service/internal/config"
	"github.com/AlexCorn999/order-data-service/internal/logger"
	"github.com/AlexCorn999/order-data-service/internal/repository"
	"github.com/AlexCorn999/order-data-service/internal/repository/postgres"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

type APIServer struct {
	config  *config.Config
	router  *chi.Mux
	logger  *log.Logger
	storage repository.Storage
}

func NewAPIServer(config *config.Config) *APIServer {
	return &APIServer{
		config: config,
		router: chi.NewRouter(),
		logger: log.New(),
	}
}

func (s *APIServer) Start(ch chan []byte) error {
	s.config.ParseFlags()
	s.configureRouter()

	if err := s.configureLogger(); err != nil {
		return err
	}

	if err := s.configureStore(); err != nil {
		return err
	}
	defer s.storage.Close()

	s.logger.Info("starting api server")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureRouter() {
	s.router.Use(logger.WithLogging)
	//s.router.Get("/order")
}

func (s *APIServer) configureLogger() error {
	level, err := log.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *APIServer) configureStore() error {
	db, err := postgres.NewStorage(s.config.DataBaseURL)
	if err != nil {
		return err
	}
	s.storage = db
	return nil
}
