package service

import (
	"encoding/json"
	"net/http"
	"strconv"

	"aviasalesTest/internal/pkg/processor"
)

func (s *Service) GetCheapest(w http.ResponseWriter, r *http.Request) {
	flights := s.processor.GetDirectionsFromDXBToBKK()

	var cheapest processor.Flight
	for _, flight := range flights {
		if cheapest.FlightNumber == "" {
			cheapest = flight
		} else {
			cheapestPrice, err := strconv.ParseFloat(cheapest.Pricing.ServiceCharges[0].Text, 64)
			if err != nil {
				s.logger.Fatalf("can not parse float price")
			}

			price, err := strconv.ParseFloat(flight.Pricing.ServiceCharges[0].Text, 64)
			if err != nil {
				s.logger.Fatalf("can not parse float price")
			}
			if price < cheapestPrice {
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
