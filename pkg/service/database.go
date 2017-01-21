package service

import (
	"github.com/codeflavor/servops/pkg/config"
)

type BackingService struct{}

func (b *BackingService) Start() error {
	return nil
}

func (b *BackingService) Status() error {
	return nil
}

func (b *BackingService) Stop() ([]Status, error) {
	return nil, nil
}

// InstatiateDBService tries to create a valid connection to the database
func InstatiateDBService(dbConfig *config.AppConfig) error {
	return nil
}
