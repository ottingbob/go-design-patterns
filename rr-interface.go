package main

import (
	"fmt"
)

type RailroadWideChecker interface {
	CheckRailsWidth() int
}

type Railroad struct {
	Width int
}

func (r *Railroad) IsCorrectSizeTrain(rrwc RailroadWideChecker) bool {
	return rrwc.CheckRailsWidth() <= r.Width
}

type Train struct {
	TrainWidth int
}

func (p *Train) CheckRailsWidth() int {
	return p.TrainWidth
}

func main() {
	railroad := Railroad{Width: 10}
	passengerTrain := &Train{TrainWidth: 10}
	cargoTrain := &Train{TrainWidth: 15}

	canPassengerTrainPass := railroad.IsCorrectSizeTrain(passengerTrain)
	canCargoTrainPass := railroad.IsCorrectSizeTrain(cargoTrain)

	fmt.Printf("Can passenger train pass? %t\n", canPassengerTrainPass)
	fmt.Printf("Can cargo train pass? %t\n", canCargoTrainPass)
}