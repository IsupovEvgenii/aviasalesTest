package service

import (
	"encoding/json"
	"net/http"

	"aviasalesTest/internal/pkg/processor"
)

func (s *Service) GetCheapest(w http.ResponseWriter, r *http.Request) {
	flights := s.processor.GetDirectionsFromDXBToBKK()

	var cheapest processor.Flight
	for _, flight := range flights {
		if cheapest.FlightNumber == "" {
			cheapest = flight
		} else {
			if s.isCheaper(flight, cheapest) {
				cheapest = flight
			}
		}
	}
	data, err := json.Marshal(cheapest)
	if err != nil {
		s.logger.Fatalf("can not unmarshal Cheapest")
	}
	w.Write(data)
}
