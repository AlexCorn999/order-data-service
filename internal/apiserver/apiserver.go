package apiserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/AlexCorn999/order-data-service/internal/config"
	"github.com/AlexCorn999/order-data-service/internal/domain"
	"github.com/AlexCorn999/order-data-service/internal/logger"
	"github.com/AlexCorn999/order-data-service/internal/nats"
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
	Nats    *nats.NatsST
}

func NewAPIServer(config *config.Config) *APIServer {
	return &APIServer{
		config: config,
		router: chi.NewRouter(),
		logger: log.New(),
	}
}

func (s *APIServer) Start(sigChan chan os.Signal) error {
	s.config.ParseFlags()
	s.configureRouter()

	if err := s.configureNats(); err != nil {
		return err
	}
	defer s.Nats.Close()

	if err := s.configureLogger(); err != nil {
		return err
	}

	if err := s.configureStore(); err != nil {
		return err
	}
	defer s.storage.Close()

	if err := s.Nats.SubscribeCh(); err != nil {
		return err
	}
	defer s.Nats.UnsubsribeNs()

	s.logger.Info("starting api server")

	go func() {
		for {
			select {
			case data := <-s.Nats.Data:

				var order domain.Order
				if err := json.Unmarshal(data, &order); err != nil {
					log.Println("JSON - something wrong")
					continue
				}
				fmt.Printf("%+v", order)

			case sig := <-sigChan:
				fmt.Println("server stoped by signal", sig)
				os.Exit(1)
			}
		}
	}()

	if err := http.ListenAndServe(s.config.BindAddr, s.router); err != nil {
		return err
	}

	return nil
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

func (s *APIServer) configureNats() error {
	nats, err := nats.NewNatsST(s.config.ClusterID, s.config.ClientID)
	if err != nil {
		return err
	}
	s.Nats = nats
	return nil
}
