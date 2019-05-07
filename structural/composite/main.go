package main

import (
	"fmt"
	
	c "composite"
)

func Swim() {
	fmt.Println("Swim var")
}

func main() {
	swimmer := c.CompositeSwimmerA{
		MySwim: Swim,
	}

	swimmer.MyAthlete.Train()
	swimmer.MySwim()

	fish := c.Shark{
		Swim: c.Swim,
	}

	fish.Eat()
	fish.Swim()

	csb := c.CompositeSwimmerB{
		&c.Athlete{},
		&c.SwimmerImpl{},
	}

	csb.Train()
	csb.Swim()
}