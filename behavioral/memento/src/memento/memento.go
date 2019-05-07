package memento

// Memento has three players called actors:

// Memento: A type that stores the type we want to save.
// Usually we wont store the business type directly and
// we provide an extra layer of abstraction through this
// type

// Originator: A type that is in charge of creating 
// mementos and storing the current active state. We said
// the Memento type wraps states of the business type and
// we use originator as the creator of mementos

// Care Taker: A type that stores the list of mementos that
// can have the logic to store them in a database or to not 
// store more than a specified number of them

// All about sequence of actions over time, to undo one or
// two operations or to provide some type of transaction
// to an application

// Captures an object state without modifying the object
// itself
// Saves a limited amount of states so we can retrieve 
// them later

import "fmt"

type Memento struct {
	State State
}

type State struct {
	Description string
}

type Originator struct {
	State State
}

func (o *Originator) NewMemento() Memento {
	return Memento{State: o.State}
}

func (o *Originator) ExtractAndStoreState(m Memento) {
	o.State = m.State
}

type CareTaker struct {
	MementoList []Memento
}

func (c *CareTaker) Add(m Memento) {
	c.MementoList = append(c.MementoList, m)
}

func (c *CareTaker) Memento(i int) (Memento, error) {
	if len(c.MementoList) < i || i < 0 {
		return Memento{}, fmt.Errorf("Index not found\n")
	}
	return c.MementoList[i], nil
}