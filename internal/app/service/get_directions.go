package service

import (
	"encoding/json"
	"net/http"
)

func (s *Service) GetDirections(w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(s.processor.GetDirectionsFromDXBToBKK())
	if err != nil {
		s.logger.Fatalf("can not unmarshal DirectionsFromDXBToBKK")
	}
	w.Write(data)
}
