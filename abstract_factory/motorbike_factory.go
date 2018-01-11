package abstract_factory

import (
	"errors"
	"fmt"
)

const (
	SportMotorbikeType  = 1
	CruiseMotorbikeType = 2
)

type MotorbikeFactory struct{}

func (m *MotorbikeFactory) NewVehicle(v int) (Vehicle, error) {
	switch v {
	case SportMotorbikeType:
		return new(SportMotorbike), nil
	case CruiseMotorbikeType:
		return new(CruiseMotorbike), nil
	default:
		return nil, errors.New(fmt.Sprintf("Vehicle of type %d not recognized\n", v))
	}
}

type SportMotorbike struct{}

func (s *SportMotorbike) GetMotorbikeType() int {
	return SportMotorbikeType
}
func (l *SportMotorbike) NumWheels() int {
	return 2
}
func (l *SportMotorbike) NumSeats() int {
	return 1
}

type CruiseMotorbike struct{}

func (s *CruiseMotorbike) GetMotorbikeType() int {
	return SportMotorbikeType
}
func (l *CruiseMotorbike) NumWheels() int {
	return 2
}
func (l *CruiseMotorbike) NumSeats() int {
	return 2
}
