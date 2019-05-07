package main

// A FSM is a pattern with one or more states and
// travels between them to execute some behaviors

// It has a type that alters its own behavior when
// some internal things have changed

// Model complex graphs and pipelines can be upgraded
// easily by adding more states and rerouting their
// output states 

import (
	"fmt"
	"os"
	"math/rand"
	"time"
)

type GameState interface {
	executeState(* GameContext) bool
}

type GameContext struct {
	SecretNumber int
	Retries int
	Won bool
	Next GameState
}

type StartState struct{}
func(s *StartState) executeState(c *GameContext) bool {
	c.Next = &AskState{}

	rand.Seed(time.Now().UnixNano())
	c.SecretNumber = rand.Intn(10)
	fmt.Println("Introduce a number or retries to set the difficulty:")
	fmt.Fscanf(os.Stdin, "%d\n", &c.Retries)

	return true
}

type FinishState struct{}
func(f *FinishState) executeState(c *GameContext) bool {
	if c.Won {
		c.Next = &WinState{}
	} else {
		c.Next = &LoseState{}
	}

	return true
}

type AskState struct{}
func (a *AskState) executeState(c *GameContext) bool {
	fmt.Printf("Introduce a number between 0 and 10, you have %d tries left\n", c.Retries)

	var n int
	fmt.Fscanf(os.Stdin, "%d", &n)
	c.Retries = c.Retries - 1

	if n == c.SecretNumber {
		c.Won = true
		c.Next = &FinishState{}
	}

	if c.Retries == 0 {
		c.Next = &FinishState{}
	}

	return true
}

type WinState struct{}

func (w *WinState) executeState(c *GameContext) bool {
	fmt.Println("Congrats, you win")
	return false
}

type LoseState struct{}

func (l *LoseState) executeState(c *GameContext) bool {
	fmt.Printf("You lose. The correct number was %d\n", c.SecretNumber)
	return false
}


func main() {
	start := StartState{}

	game := GameContext{
		Next: &start,
	}

	for game.Next.executeState(&game) {}
}