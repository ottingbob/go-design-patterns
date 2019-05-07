package singleton

import (
	"testing"

	"singleton"
)

func TestGetInstance(t *testing.T) {
	counter1 := singleton.GetInstance()

	if counter1 == nil {
		// Test acceptance criteria 1 failed
		t.Error("Expected pointer to Singleton after calling GetInstance(), not nil")
	}

	expectedCounter := counter1

	currentCount := counter1.AddOne()
	if currentCount != 1 {
		t.Errorf("After calling for the first time to count, the count must be 1 " +
			"but it is %d\n", currentCount)
	}

	counter2 := singleton.GetInstance()
	if counter2 != expectedCounter {
		// Test 2 failed
		t.Error("Expected same instance in counter2 but it got a different instance")
	}

	currentCount = counter2.AddOne()
	if currentCount != 2 {
		t.Errorf("After calling 'AddOne' using the second counter, the current count " +
			"must be 2 but it was %d\n", currentCount)
	}
}