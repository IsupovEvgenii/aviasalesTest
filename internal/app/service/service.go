package service

import (
	"log"

	"aviasalesTest/internal/pkg/processor"
)

type Processor interface {
	GetDirectionsFromDXBToBKK() []processor.Flight
}

type Service struct {
	logger *log.Logger
	processor Processor
}

func NewService(l *log.Logger, p Processor) *Service {
	return &Service{logger: l, processor: p}
}
