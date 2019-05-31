package app

import (
	"log"
	"os"
)

const (
	prefixAviaSales = "aviasalesTest "
)
type Application struct {
	Logger *log.Logger
}

func New() *Application {
	return &Application{Logger: log.New(os.Stdout, prefixAviaSales, 0)}
}

func (a *Application) Init() error {
	return nil
}