package processor

import (
	"log"

	"aviasalesTest/internal/pkg/apiresolver"
)

type APIResolver interface {
	GetFirstResponse() apiresolver.AirFareSearchResponse
	GetSecondResponse() apiresolver.AirFareSearchResponse
}

// Service implements service
type Service struct {
	logger   *log.Logger
	resolver APIResolver
}

// NewService creates new service
func NewService(l *log.Logger, ar APIResolver) *Service {
	return &Service{logger: l, resolver: ar}
}

type Flight struct {
	Carrier struct {
		Text string `json:",chardata"`
		ID   string `json:"id,attr"`
	} `json:"Carrier"`
	FlightNumber       string `json:"FlightNumber"`
	Source             string `json:"Source"`
	Destination        string `json:"Destination"`
	DepartureTimeStamp string `json:"DepartureTimeStamp"`
	ArrivalTimeStamp   string `json:"ArrivalTimeStamp"`
	Class              string `json:"Class"`
	NumberOfStops      string `json:"NumberOfStops"`
	FareBasis          string `json:"FareBasis"`
	WarningText        string `json:"WarningText"`
	TicketType         string `json:"TicketType"`
	Pricing            struct {
		Text           string           `json:",chardata"`
		Currency       string           `json:"currency,attr"`
		ServiceCharges []ServicesCharge `json:"ServiceCharges"`
	} `json:"Pricing"`
	TypePricedItinerary string `json:"TypePricedItinerary"`
}

type ServicesCharge struct {
	Text       string `json:",chardata"`
	Type       string `json:"type,attr"`
	ChargeType string `json:"ChargeType,attr"`
}

const (
	from       = "DXB"
	to         = "BKK"
	TypeOnward = "onward"
	TypeReturn = "return"
)

func (s *Service) GetDirectionsFromDXBToBKK() []Flight {
	var responses []apiresolver.AirFareSearchResponse
	responses = append(responses, s.resolver.GetFirstResponse())
	responses = append(responses, s.resolver.GetSecondResponse())

	var resultFlights []Flight
	for _, resp := range responses {
		for _, pricedItineraryFlight := range resp.PricedItineraries.Flights {
			for _, onwardFlight := range pricedItineraryFlight.OnwardPricedItinerary.Flights.Flight {
				if onwardFlight.Source == from && onwardFlight.Destination == to {
					curFlight := Flight{
						FlightNumber:        onwardFlight.FlightNumber,
						Source:              onwardFlight.Source,
						Destination:         onwardFlight.Destination,
						DepartureTimeStamp:  onwardFlight.DepartureTimeStamp,
						ArrivalTimeStamp:    onwardFlight.ArrivalTimeStamp,
						Class:               onwardFlight.Class,
						NumberOfStops:       onwardFlight.NumberOfStops,
						FareBasis:           onwardFlight.FareBasis,
						WarningText:         onwardFlight.WarningText,
						TicketType:          onwardFlight.TicketType,
						TypePricedItinerary: TypeOnward,
					}
					curFlight.Carrier.Text = onwardFlight.Carrier.Text
					curFlight.Carrier.ID = onwardFlight.Carrier.ID
					curFlight.Pricing.Text = pricedItineraryFlight.Pricing.Text
					curFlight.Pricing.Currency = pricedItineraryFlight.Pricing.Currency
					curFlight.Pricing.ServiceCharges = make([]ServicesCharge, len(pricedItineraryFlight.Pricing.ServiceCharges))
					for i, serviceCharge := range pricedItineraryFlight.Pricing.ServiceCharges {
						curFlight.Pricing.ServiceCharges[i].Text = serviceCharge.Text
						curFlight.Pricing.ServiceCharges[i].Type = serviceCharge.Type
						curFlight.Pricing.ServiceCharges[i].ChargeType = serviceCharge.ChargeType
					}

					resultFlights = append(resultFlights, curFlight)
				}
			}

			for _, returnFlight := range pricedItineraryFlight.ReturnPricedItinerary.Flights.Flight {
				if returnFlight.Source == from && returnFlight.Destination == to {
					curFlight := Flight{
						FlightNumber:        returnFlight.FlightNumber,
						Source:              returnFlight.Source,
						Destination:         returnFlight.Destination,
						DepartureTimeStamp:  returnFlight.DepartureTimeStamp,
						ArrivalTimeStamp:    returnFlight.ArrivalTimeStamp,
						Class:               returnFlight.Class,
						NumberOfStops:       returnFlight.NumberOfStops,
						FareBasis:           returnFlight.FareBasis,
						WarningText:         returnFlight.WarningText,
						TicketType:          returnFlight.TicketType,
						TypePricedItinerary: TypeReturn,
					}
					curFlight.Carrier.Text = returnFlight.Carrier.Text
					curFlight.Carrier.ID = returnFlight.Carrier.ID
					curFlight.Pricing.Text = pricedItineraryFlight.Pricing.Text
					curFlight.Pricing.Currency = pricedItineraryFlight.Pricing.Currency
					curFlight.Pricing.ServiceCharges = make([]ServicesCharge, len(pricedItineraryFlight.Pricing.ServiceCharges))
					for i, serviceCharge := range pricedItineraryFlight.Pricing.ServiceCharges {
						curFlight.Pricing.ServiceCharges[i].Text = serviceCharge.Text
						curFlight.Pricing.ServiceCharges[i].Type = serviceCharge.Type
						curFlight.Pricing.ServiceCharges[i].ChargeType = serviceCharge.ChargeType
					}
					resultFlights = append(resultFlights, curFlight)
				}
			}
		}
	}

	return resultFlights
}
