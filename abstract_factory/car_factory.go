package abstract_factory

import (
	"errors"
	"fmt"
)

const (
	LuxuryCarType = 1
	FamilyCarType = 2
)

type CarFactory struct{}

func (c *CarFactory) NewVehicle(v int) (Vehicle, error) {
	switch v {
	case LuxuryCarType:
		return new(LuxuryCar), nil
	case FamilyCarType:
		return new(FamilyCar), nil
	default:
		return nil, errors.New(fmt.Sprintf("Vehicle of type %d not recognizes\n", v))
	}
}

type LuxuryCar struct{}

func (l *LuxuryCar) NumDoors() int {
	return 4
}
func (l *LuxuryCar) NumWheels() int {
	return 4
}
func (l *LuxuryCar) NumSeats() int {
	return 5
}

type FamilyCar struct{}

func (f *FamilyCar) NumDoors() int {
	return 5
}
func (l *FamilyCar) NumWheels() int {
	return 4
}
func (l *FamilyCar) NumSeats() int {
	return 5
}
