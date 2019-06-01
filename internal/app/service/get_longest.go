package service

import (
	"encoding/json"
	"net/http"
	"time"

	"aviasalesTest/internal/pkg/processor"
)

func (s *Service) GetLongest(w http.ResponseWriter, r *http.Request) {
	flights := s.processor.GetDirectionsFromDXBToBKK()

	locMoscow, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		s.logger.Fatalf("can not load location asia/bangkok")
	}

	var result processor.Flight
	for _, flight := range flights {
		if result.FlightNumber == "" {
			result = flight
		} else {
			resultTimeArrival, err := time.Parse("2006-01-02T1504", result.ArrivalTimeStamp)
			if err != nil {
				s.logger.Fatalf("can not parse time arrival")
			}

			resultTimeDeparture, err := time.Parse("2006-01-02T1504", result.DepartureTimeStamp)
			if err != nil {
				s.logger.Fatalf("can not parse time departure")
			}

			resultDelta := resultTimeArrival.In(locMoscow).Sub(resultTimeDeparture.In(locMoscow))

			curTimeArrival, err := time.Parse("2006-01-02T1504", flight.ArrivalTimeStamp)
			if err != nil {
				s.logger.Fatalf("can not parse time cur arrival")
			}

			curTimeDeparture, err := time.Parse("2006-01-02T1504", flight.DepartureTimeStamp)
			if err != nil {
				s.logger.Fatalf("can not parse time cur depurture")
			}

			curDelta := curTimeArrival.In(locMoscow).Sub(curTimeDeparture.In(locMoscow))

			if curDelta > resultDelta {
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
