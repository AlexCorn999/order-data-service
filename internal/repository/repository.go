package repository

type Storage interface {
	Close() error
}
