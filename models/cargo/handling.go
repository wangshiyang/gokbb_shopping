package cargo

import (
	"errors"
	"gokbb_shopping/models/location"
	"time"
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
	VoyageNumber Number
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

type HandlingEventRepository interface {
	Store(e HandlingEvent)
	QueryHandlingHistory(TrackingID) HandlingHistory
}

type HandlingEventFactory struct {
	CargoRepository    CargoRepository
	VoyageRepository   VoyageRepository
	LocationRepository location.LocationRepository
}

func (h HandlingHistory) MostRecentlyCompletedEvent() (HandlingEvent, error) {
	if len(h.HandlingEvent) == 0 {
		return HandlingEvent{}, errors.New("delivery history is empty")
	}

	return h.HandlingEvent[len(h.HandlingEvent)-1], nil
}

func (f *HandlingEventFactory) CreateHandlingEvent(registered time.Time, completed time.Time, id TrackingID,
	voyageId Number, unlocode location.UNLocode, eventType HandlingTypeEvent) (HandlingEvent, error) {
	if _, err := f.CargoRepository.Find(id); err != nil {
		return HandlingEvent{}, err
	}

	if _, err := f.VoyageRepository.Find(voyageId); err != nil {
		if len(voyageId) > 0 {
			return HandlingEvent{}, err
		}
	}

	if _, err := f.LocationRepository.Find(unlocode); err != nil {
		return HandlingEvent{}, err
	}

	return HandlingEvent{
		TrackingID: id,
		Activity: HandlingActivity{
			Type:         eventType,
			Location:     unlocode,
			VoyageNumber: voyageId,
		},
	}, nil

}
