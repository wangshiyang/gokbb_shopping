package cargo

import (
	"shawn/gokbb_shopping/models/location"
	"github.com/go-kit/kit/examples/shipping/voyage"
	"errors"
)

const (
	NotHandled HandlingTypeEvent = iota
	Load
	Unload
	Receive
	Claim
	Customs
)

type HandlingTypeEvent int

type HandlingEvent struct {
	TrackingID TrackingID
	Activity   HandlingActivity
}

type HandlingActivity struct {
	Type         HandlingTypeEvent
	Location     location.UNLocode
	VoyageNumber voyage.Number
}

func (t HandlingTypeEvent) String() string {
	switch t {
	case NotHandled:
		return "Not Handled"
	case Load:
		return "Load"
	case Unload:
		return "Unload"
	case Receive:
		return "Receive"
	case Claim:
		return "Claim"
	case Customs:
		return "Customs"
	}

	return ""
}

type HandlingHistory struct {
	HandlingEvent []HandlingEvent
}

func (h HandlingHistory) MostRecentlyCompletedEvent() (HandlingEvent, error) {
	if len(h.HandlingEvent) == 0 {
		return HandlingEvent{}, errors.New("delivery history is empty")
	}

	return h.HandlingEvent[len(h.HandlingEvent)-1], nil
}


