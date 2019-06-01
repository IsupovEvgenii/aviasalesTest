package service

import (
	"encoding/json"
	"net/http"

	"aviasalesTest/internal/pkg/processor"
)

func (s *Service) GetMostExpensive(w http.ResponseWriter, r *http.Request) {
	flights := s.processor.GetDirectionsFromDXBToBKK()

	var expensive processor.Flight
	for _, flight := range flights {
		if expensive.FlightNumber == "" {
			expensive = flight
		} else {
			if !s.isCheaper(flight, expensive) {
				expensive = flight
			}
		}

	}
	data, err := json.Marshal(expensive)
	if err != nil {
		s.logger.Fatalf("can not unmarshal expensive")
	}
	w.Write(data)
}
