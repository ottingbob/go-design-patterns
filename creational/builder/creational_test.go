
// Should have named this builder...
package creational

import (
	"testing"

	"creational"
)

func TestBuilderPattern(t *testing.T) {
	manufacturingComplex := creational.ManufacturingDirector{}

	carBuilder := &creational.CarBuilder{}

	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()

	car := carBuilder.GetVehicle()

	if car.Wheels != 4 {
		t.Errorf("Wheels on a car must be 4 and they were %d\n", car.Wheels)
	}

	if car.Structure != "Car" {
		t.Errorf("Structure on a car must be 'Car' and was %s\n", car.Structure)
	}

	if car.Seats != 5 {
		t.Errorf("Seats on a car must be 5 and they were %d\n", car.Seats)
	}

	bikeBuilder := &creational.BikeBuilder{}

	manufacturingComplex.SetBuilder(bikeBuilder)
	manufacturingComplex.Construct()

	motorbike := bikeBuilder.GetVehicle()
	motorbike.Seats = 1

	if motorbike.Wheels != 2 {
		t.Errorf("Wheels on a motobike must be 2 and they were %d\n", motorbike.Wheels)
	}

	if motorbike.Structure != "Motorbike" {
		t.Errorf("Structure on a motorbike must be 'Motorbike' and was %s\n", motorbike.Structure)
	}



}