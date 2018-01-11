package abstract_factory

import (
	"testing"
)

func TestCarFactory(t *testing.T) {
	carF, err := BuildFactory(CarFactoryType)
	if err != nil {
		t.Fatal(err)
	}
	carVehicle, err := carF.NewVehicle(LuxuryCarType)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Car vehicle has %d seats.", carVehicle.NumSeats())

	luxuryCar, ok := carVehicle.(Car)
	if !ok {
		t.Fatal("Struct assertion failed.")
	}
	t.Logf("Luxury car has %d doors.", luxuryCar.NumDoors())
}
