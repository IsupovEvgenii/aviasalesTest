package service

import (
	"encoding/json"
	"net/http"
	"strconv"

	"aviasalesTest/internal/pkg/processor"
)

func (s *Service) GetMostExpensive(w http.ResponseWriter, r *http.Request) {
	flights := s.processor.GetDirectionsFromDXBToBKK()

	var expensive processor.Flight
	for _, flight := range flights {
		if expensive.FlightNumber == "" {
			expensive = flight
		} else {
			expensivePrice, err := strconv.ParseFloat(expensive.Pricing.ServiceCharges[0].Text, 64)
			if err != nil {
				s.logger.Fatalf("can not parse float price")
			}

			price, err := strconv.ParseFloat(flight.Pricing.ServiceCharges[0].Text, 64)
			if err != nil {
				s.logger.Fatalf("can not parse float price")
			}
			if price > expensivePrice {
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
