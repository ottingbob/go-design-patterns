package memento

import (
	"testing"
	m "memento"
)

func TestCareTaker_Add(t *testing.T) {
	originator := m.Originator{}
	originator.State = m.State{Description: "Idle"}

	careTaker := m.CareTaker{}

	mem := originator.NewMemento()
	if mem.State.Description != "Idle" {
		t.Error("Expected state was not found")
	}

	currentLen := len(careTaker.MementoList)
	careTaker.Add(mem)

	if len(careTaker.MementoList) != currentLen + 1 {
		t.Error("No new elements were added on the list")
	}
}

func TestCareTaker_Memento(t *testing.T) {
	originator := m.Originator{}
	careTaker := m.CareTaker{}

	originator.State = m.State{"Idle"}
	careTaker.Add(originator.NewMemento())

	mem, err := careTaker.Memento(0)
	if err != nil {
		t.Fatal(err)
	}

	if mem.State.Description != "Idle" {
		t.Error("Unexpected state")
	}

	mem, err = careTaker.Memento(-1)
	if err == nil {
		t.Fatal("An error is expected when asking for a negative number but no error was found")
	}
}

func TestOriginator_ExtractAndStoreState(t *testing.T) {
	originator := m.Originator{State: m.State{"Idle"}}
	idleMemento := originator.NewMemento()

	originator.State = m.State{"Working"}

	originator.ExtractAndStoreState(idleMemento)
	if originator.State.Description != "Idle" {
		t.Error("Unexpected state found")
	}
}