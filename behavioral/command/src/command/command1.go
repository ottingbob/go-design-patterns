package main

// Focus on invocation of something or on the abstraction
// of some type

// Commonly seen as a container. You put something like
// the info for user interaction on a UI and pass a command.
// Therefore you focus on calling the command as opposed
// to the implementation of it

// Similar to strategy in that the interface may have different
// commands pointing to different algorithms however the implementation
// is actually housed somewhere else

import (
	"fmt"
	// "net/http"
)

type Command interface {
	Execute()
}

type ConsoleOutput struct {
	message string
}

func (c *ConsoleOutput) Execute() {
	fmt.Println(c.message)
}

func CreateCommand(s string) Command {
	fmt.Println("Creating command")

	return &ConsoleOutput{
		message: s,
	}
}

type CommandQueue struct {
	queue []Command
}

func (p *CommandQueue) AddCommand(c Command) {
	p.queue = append(p.queue, c)

	if len(p.queue) == 3 {
		for _, command := range p.queue {
			command.Execute()
		}

		p.queue = make([]Command, 3)
	}
}

func main() {
	queue := CommandQueue{}

	queue.AddCommand(CreateCommand("First message"))
	queue.AddCommand(CreateCommand("Second message"))
	queue.AddCommand(CreateCommand("Third message"))

	queue.AddCommand(CreateCommand("Fourth message"))
	queue.AddCommand(CreateCommand("Fifth message"))

	// client := http.Client{}
	// client.Do(nil)
}