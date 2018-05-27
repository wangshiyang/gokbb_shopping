package location

import (
	"github.com/go-kit/kit/examples/shipping/location"
	"time"
)

type Leg struct {
	LoadLocation   location.UNLocode
	UnLoadLocation location.UNLocode
	LoadTime       time.Time
	UnLoadTime     time.Time
}

type Itinerary struct {
	Legs []Leg `json:"legs"`
}

func (i Itinerary) InitialDepartureLocation() location.UNLocode {
	if i.IsEmpty(){
		return  location.UNLocode("")
	}

	return i.Legs[0].LoadLocation
}

func (i Itinerary) FinalArrivalLocation() location.UNLocode {
	if i.IsEmpty() {
		return location.UNLocode("")
	}

	return i.Legs[0].UnLoadLocation
}

func (i Itinerary)FinalArrivalTime() time.Time {
	return i.Legs[0].UnLoadTime
}

func (i Itinerary) IsEmpty() bool {
	return i.Legs == nil || len(i.Legs) == 0
}

func (i Itinerary) IsExpected() bool {
	return true
}
