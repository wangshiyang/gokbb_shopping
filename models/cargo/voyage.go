package cargo

import (
	"shawn/gokbb_shopping/models/location"
	"time"
	"errors"
)

type Number string

type Voyage struct {
	Number   Number
	Schedule Schedule
}

type Schedule struct {
	CarrierMovement []CarrierMovement
}

type CarrierMovement struct {
	DepartureLocation location.UNLocode
	ArrivalLocation   location.UNLocode
	DepartureTime     time.Time
	ArrivalTime       time.Time
}

var ErrUnknown = errors.New("unknow voyage")

func New(n Number, s Schedule) *Voyage {
	return &Voyage{Number: n, Schedule: s}
}

type VoyageRepository interface {
	Find(Number) (*Voyage, error)
}
