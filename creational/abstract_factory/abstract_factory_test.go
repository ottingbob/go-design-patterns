package abstract_factory

import (
	"testing"

	af "abstract_factory"
)

func TestMotorbikeFactory(t *testing.T) {
	motorbikeF, err := af.BuildFactory(af.MotorbikeFactoryType)
	if (err != nil) {
		t.Fatal(err)
	}

	motorbikeVehicle, err := motorbikeF.NewVehicle(af.SportMotorbikeType)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Motorbike vehicle has %d wheels\n", motorbikeVehicle.NumWheels())

	sportBike, ok := motorbikeVehicle.(af.Motorbike)
	if !ok {
		t.Log("LOG:", ok)
		t.Fatal("Struct assertion has failed")
	}

	t.Logf("Sport motorbike has type %d\n", sportBike.GetMotorbikeType())
} 

func TestCarFactory(t *testing.T) {
	carF, err := af.BuildFactory(af.CarFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	carVehicle, err := carF.NewVehicle(af.LuxuryCarType)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Car vehicle has %d seats\n", carVehicle.NumWheels())

	luxuryCar, ok := carVehicle.(af.Car)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}
	t.Logf("Luxury car has %d doors.\n", luxuryCar.NumDoors())
}