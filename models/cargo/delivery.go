package cargo

import (
	"shawn/gokbb_shopping/models"
	"shawn/gokbb_shopping/models/location"
)

type Delivery struct {
	models.Model
	Itinerary location.Itinerary
	RouteSpecification RouteSpecification

}
