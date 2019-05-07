package observer

// Also known as the publish/subscriber or
// publish/listener pattern

// Pattern aims to subscribe to some event that will
// trigger some behavior on many subscribed types
// As a result we uncouple an event from its possible
// handler functions

// Useful to achieve many actions that are triggered
// on one event or when you don't know how many actions
// are performed after an event in advance or that the
// number of actions are going to grow

// Provides an event driven architecture where on event
// can trigger one or more actions
// Uncouple the actions that are performed from the event
// that triggers them
// Provide more than one event that triggers the same action

import "fmt"

type Observer interface {
	Notify(string)
}

type Publisher struct {
	ObserversList []Observer
}

func (s *Publisher) AddObserver(o Observer) {
	s.ObserversList = append(s.ObserversList, o)
}

func (s *Publisher) RemoveObserver(o Observer) {
	var indexToRemove int

	for i, observer := range s.ObserversList {
		if observer == o {
			indexToRemove = i
			break
		}
	}

	s.ObserversList = append(s.ObserversList[:indexToRemove], s.ObserversList[indexToRemove+1:]...)
}

func (s *Publisher) NotifyObservers(m string) {
	fmt.Printf("Publisher received message '%s' to notify observers\n", m)
	for _, observer := range s.ObserversList {
		observer.Notify(m)
	}
}