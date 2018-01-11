package builder

import (
	"strings"
	"testing"
)

func TestCarBuild(t *testing.T) {
	manufacturingComplex := ManufactoringDirector{}
	carBuilder := &CarBuilder{}
	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct(carBuilder)
	car := carBuilder.Build()

	if car.Wheels != 4 {
		t.Errorf("Wheels on a car must be 5 and they were: %d\n", car.Wheels)
	}
	if strings.ToLower(car.Structure) != "car" {
		t.Errorf("Structure on a car must be 'Car'")
	}
	if car.Seats != 5 {
		t.Errorf("Seats on a car must bne 5 and they were %d\n", car.Seats)
	}
}

func TestBikeBuild(t *testing.T) {
	manufacturingComplex := ManufactoringDirector{}
	bikeBuilder := &BikeBuilder{}
	manufacturingComplex.SetBuilder(bikeBuilder)
	manufacturingComplex.Construct(bikeBuilder)
	bike := bikeBuilder.Build()

	if bike.Wheels != 2 {
		t.Errorf("Wheels on a bike must be 2 and they were: %d\n", bike.Wheels)
	}
	if strings.ToLower(bike.Structure) != "bike" {
		t.Errorf("Structure on a bike must be 'Car'")
	}
	if bike.Seats != 2 {
		t.Errorf("Seats on a bike must bne 5 and they were %d\n", bike.Seats)
	}
}
