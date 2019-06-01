package service

import (
	"strconv"
	"time"

	"aviasalesTest/internal/pkg/processor"
)

func (s *Service) isLonger(flight1 processor.Flight, flight2 processor.Flight) bool {
	locMoscow, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		s.logger.Fatalf("can not load location asia/bangkok")
	}

	resultTimeArrival, err := time.Parse("2006-01-02T1504", flight1.ArrivalTimeStamp)
	if err != nil {
		s.logger.Fatalf("can not parse time arrival")
	}

	resultTimeDeparture, err := time.Parse("2006-01-02T1504", flight1.DepartureTimeStamp)
	if err != nil {
		s.logger.Fatalf("can not parse time departure")
	}

	resultDelta := resultTimeArrival.In(locMoscow).Sub(resultTimeDeparture.In(locMoscow))

	curTimeArrival, err := time.Parse("2006-01-02T1504", flight2.ArrivalTimeStamp)
	if err != nil {
		s.logger.Fatalf("can not parse time cur arrival")
	}

	curTimeDeparture, err := time.Parse("2006-01-02T1504", flight2.DepartureTimeStamp)
	if err != nil {
		s.logger.Fatalf("can not parse time cur depurture")
	}

	curDelta := curTimeArrival.In(locMoscow).Sub(curTimeDeparture.In(locMoscow))

	return resultDelta >= curDelta
}

func (s *Service) isCheaper(flight1 processor.Flight, flight2 processor.Flight) bool {
	cheapestPrice, err := strconv.ParseFloat(flight1.Pricing.ServiceCharges[0].Text, 64)
	if err != nil {
		s.logger.Fatalf("can not parse float price")
	}

	price, err := strconv.ParseFloat(flight2.Pricing.ServiceCharges[0].Text, 64)
	if err != nil {
		s.logger.Fatalf("can not parse float price")
	}
	return cheapestPrice <= price
}
