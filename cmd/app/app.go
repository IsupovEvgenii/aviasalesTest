package app

import (
	"aviasalesTest/internal/pkg/processor"
	"log"
	"os"

	"aviasalesTest/internal/app/service"
	"aviasalesTest/internal/pkg/apiresolver"
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

	resolver := apiresolver.NewService(a.Logger, "apiresponses/RS_Via-3.xml", "apiresponses/RS_ViaOW.xml")
	processor := processor.NewService(a.Logger, resolver)
	service := service.NewService(a.Logger, processor)

	a.Service = service
	return nil
}