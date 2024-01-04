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
	"github.com/AlexCorn999/order-data-service/internal/repository/inMemory"
	"github.com/AlexCorn999/order-data-service/internal/repository/postgres"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

type APIServer struct {
	config       *config.Config
	router       *chi.Mux
	logger       *log.Logger
	postgres     repository.Storage
	cacheStorage repository.Storage
	nats         *nats.NatsST
}

func NewAPIServer(config *config.Config) *APIServer {
	return &APIServer{
		config: config,
		router: chi.NewRouter(),
		logger: log.New(),
	}
}

// Start starts the server, configures routing, parses data from the config, configures nats and opens the database connection.
func (s *APIServer) Start(sigChan chan os.Signal) error {
	s.config.ParseFlags()
	s.configureRouter()

	if err := s.configureNats(); err != nil {
		return err
	}
	defer s.nats.Close()

	if err := s.configureLogger(); err != nil {
		return err
	}

	if err := s.configureStore(); err != nil {
		return err
	}
	defer s.postgres.Close()

	if err := s.nats.SubscribeCh(); err != nil {
		return err
	}
	defer s.nats.UnsubsribeNs()

	s.logger.Info("starting api server")

	// добавить восстановление данных из бд в кэш

	// reading data from the channel, validation and writing to the database. Application termination by signal.
	go func() {
		for {
			select {
			case data := <-s.nats.Data:

				var order domain.Order
				if err := json.Unmarshal(data, &order); err != nil {
					log.Println("JSON - something wrong")
					continue
				}

				if err := order.Validate(); err != nil {
					log.Println("incorrect data has been received: ", err)
					continue
				}

				// добавить валидацию и запись в хранилище + кэш
				fmt.Printf("%+v\n", order)

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

// configureRouter configures routing for requests.
func (s *APIServer) configureRouter() {
	s.router.Use(logger.WithLogging)
	//s.router.Get("/order")
}

// configureLogger configures the logger for operation and specifies the logging level.
func (s *APIServer) configureLogger() error {
	level, err := log.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

// configureStore configures the database connection.
func (s *APIServer) configureStore() error {
	db, err := postgres.NewStorage(s.config.DataBaseURL)
	if err != nil {
		return err
	}

	s.cacheStorage = inMemory.NewStorage()
	s.postgres = db
	return nil
}

// configureNats connects to nats.
func (s *APIServer) configureNats() error {
	nats, err := nats.NewNatsST(s.config.ClusterID, s.config.ClientID)
	if err != nil {
		return err
	}
	s.nats = nats
	return nil
}
