package cargo

import (
	"shawn/gokbb_shopping/models"
	"shawn/gokbb_shopping/models/location"
	"time"
)

// TrackingID 是 cargo表的 唯一主键
type TrackingID string

type Cargo struct {
	models.Model


}

type RouteSpecification struct {
	Origin location.UNLocode
	Destination location.UNLocode
	ArrivalDeadline time.Time
}

func (s RouteSpecification) IsSatisfieldBy()  {
	
}

