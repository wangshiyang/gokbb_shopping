package cargo

import (
	"gokbb_shopping/models"
	"gokbb_shopping/models/location"
)

type Delivery struct {
	models.Model
	Itinerary          location.Itinerary
	RouteSpecification RouteSpecification
}
