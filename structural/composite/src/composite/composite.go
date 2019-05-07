package composite

import (
	"fmt"
)

// direct composition is having everyting we need
// as fields within the struct
//

type Athlete struct{}

func (a *Athlete) Train() {
	fmt.Println("Training")
}

type CompositeSwimmerA struct {
	MyAthlete Athlete
	MySwim func()
}

type Swimmer interface {
	Swim()
}

type Trainer interface {
	Train()
}

type SwimmerImpl struct{}
func (s *SwimmerImpl) Swim() {
	println("Swimming!")
}

type CompositeSwimmerB struct{
	Trainer
	Swimmer
}

func Swim() {
	fmt.Println("Swimming!")
}

type Animal struct{}

func (r *Animal)Eat() {
	println("Eating")
}

type Shark struct {
	Animal
	Swim func()
}
