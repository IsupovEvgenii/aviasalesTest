package service

import "log"

type Service struct {
	logger *log.Logger
}

func NewService(l *log.Logger) *Service {
	return &Service{logger: l}
}
