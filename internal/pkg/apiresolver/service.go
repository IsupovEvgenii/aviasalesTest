package apiresolver

import (
	"encoding/xml"
	"io/ioutil"
	"log"
)

// Service implements service
type Service struct {
	logger    *log.Logger
	responses []AirFareSearchResponse
}

// NewService creates new service
func NewService(l *log.Logger, fileName1 string, fileName2 string) *Service {
	data1, err := ioutil.ReadFile(fileName1)
	if err != nil {
		l.Fatalf("can not read file1")
	}

	var resp AirFareSearchResponse
	var responses []AirFareSearchResponse
	err = xml.Unmarshal(data1, &resp)
	if err != nil {
		l.Fatalf("can not unmarshal data1")
	}
	responses = append(responses, resp)

	data2, err := ioutil.ReadFile(fileName2)
	if err != nil {
		l.Fatalf("can not read file2")
	}

	err = xml.Unmarshal(data2, &resp)
	if err != nil {
		l.Fatalf("can not unmarshal data2")
	}
	responses = append(responses, resp)

	return &Service{logger: l, responses: responses}
}

// GetFirstResponse gets first response
func (s *Service) GetFirstResponse() AirFareSearchResponse {
	return s.responses[0]
}

// GetSecondResponse gets second response
func (s *Service) GetSecondResponse() AirFareSearchResponse {
	return s.responses[1]
}
