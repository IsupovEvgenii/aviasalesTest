package app

import (
	"log"
	"os"

	"aviasalesTest/internal/app/service"
)

const (
	prefixAviaSales = "aviasalesTest "
)
type Application struct {
	Logger *log.Logger
	Service *service.Service
}

func New() *Application {
	return &Application{Logger: log.New(os.Stdout, prefixAviaSales, 0)}
}

func (a *Application) Init() error {

	service := service.NewService(a.Logger)

	a.Service = service
	return nil
}