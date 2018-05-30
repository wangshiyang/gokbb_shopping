package location

import "github.com/pkg/errors"

type UNLocode string

type Location struct {
	UNLocode UNLocode
	Name     string
}

var ErrUnknown = errors.New("unknown location")

type LocationRepository interface {
	Find(locode UNLocode) (*Location, error)
	FindAll() []*Location
}
