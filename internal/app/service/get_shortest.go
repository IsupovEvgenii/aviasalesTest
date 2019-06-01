package service

import (
	"encoding/json"
	"net/http"

	"aviasalesTest/internal/pkg/processor"
)

func (s *Service) GetShortest(w http.ResponseWriter, r *http.Request) {
	flights := s.processor.GetDirectionsFromDXBToBKK()

	var result processor.Flight
	for _, flight := range flights {
		if result.FlightNumber == "" {
			result = flight
		} else {
			if !s.isLonger(flight, result) {
				result = flight
			}
		}
	}
	data, err := json.Marshal(result)
	if err != nil {
		s.logger.Fatalf("can not unmarshal result")
	}
	w.Write(data)
}
