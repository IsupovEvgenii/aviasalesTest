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
func NewService(l *log.Logger, fileNames []string) *Service {
	var responses []AirFareSearchResponse
	for _, fileName := range fileNames {
		data, err := ioutil.ReadFile(fileName)
		if err != nil {
			l.Fatalf("can not read file1")
		}

		var resp AirFareSearchResponse
		err = xml.Unmarshal(data, &resp)
		if err != nil {
			l.Fatalf("can not unmarshal data1")
		}
		responses = append(responses, resp)
	}

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
