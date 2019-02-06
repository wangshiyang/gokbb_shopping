package cargo

import (
	"gokbb_shopping/models"
	"gokbb_shopping/models/location"
	"time"
)

// TrackingID 是 cargo表的 唯一主键
type TrackingID string

type Cargo struct {
	models.Model
	Origin             location.UNLocode
	RouteSpecification RouteSpecification
	Itinerary          location.Itinerary
	Delivery           Delivery
}

func (c *Cargo) SpecifyNewRoute(rs RouteSpecification) {
	c.RouteSpecification = rs
}

type RouteSpecification struct {
	Origin          location.UNLocode
	Destination     location.UNLocode
	ArrivalDeadline time.Time
}

func (s RouteSpecification) IsSatisfieldBy(itinerary location.Itinerary) bool {
	return itinerary.Legs != nil &&
		s.Origin == itinerary.InitialDepartureLocation() &&
		s.Destination == itinerary.FinalArrivalLocation()
}

type RoutingStatus int

const (
	NotRouted RoutingStatus = iota
	Misrouted
	Routed
)

func (s RoutingStatus) String() string {
	switch s {
	case NotRouted:
		return "Not routed"
	case Misrouted:
		return "Misrouted"
	case Routed:
		return "Routed"
	}

	return ""
}

type TransprotStatus int

const (
	NotReceived TransprotStatus = iota
	InPort
	OnboardCarrier
	Claimed
	Unknown
)

type CargoRepository interface {
	Store(cargo *Cargo) error
	Find(id TrackingID) (*Cargo, error)
	FindAll() []*Cargo
}

func (t TransprotStatus) String() string {
	switch t {
	case NotReceived:
		return "Not Received"
	case InPort:
		return "InPort"
	case OnboardCarrier:
		return "Onboard Carrier"
	case Claimed:
		return "Claimed"
	case Unknown:
		return "Unknown"
	}
}
