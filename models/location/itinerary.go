package location

import (
	"time"
)

type Leg struct {
	LoadLocation   UNLocode
	UnLoadLocation UNLocode
	LoadTime       time.Time
	UnLoadTime     time.Time
}

type Itinerary struct {
	Legs []Leg `json:"legs"`
}

func (i Itinerary) InitialDepartureLocation() UNLocode {
	if i.IsEmpty(){
		return  UNLocode("")
	}

	return i.Legs[0].LoadLocation
}

func (i Itinerary) FinalArrivalLocation() UNLocode {
	if i.IsEmpty() {
		return UNLocode("")
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
