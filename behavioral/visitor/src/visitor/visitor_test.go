package visitor 

import (
	"testing"

	v "visitor"
)

type TestHelper struct {
	Received string
}

func (t *TestHelper) Write(p []byte) (int, error) {
	t.Received = string(p)
	return len(p), nil
}

func Test_Overall(t *testing.T) {
	testHelper := &TestHelper{}
	visitor := &v.MessageVisitor{}

	t.Run("MessageA test", func(t *testing.T){
		msg := v.MessageA{
			Msg: "Hello World",
			Output: testHelper,
		}

		msg.Accept(visitor)
		msg.Print()

		expected := "A: Hello World (Visited A)"
		if testHelper.Received != expected {
			t.Errorf("Expected result was incorrect. %s != %s",
				testHelper.Received, expected)
		}
	})

	t.Run("MessageB test", func(t *testing.T) {
		msg := v.MessageB{
			Msg: "Hello World",
			Output: testHelper,
		}

		msg.Accept(visitor)
		msg.Print()

		expected := "B: Hello World (Visited B)"
		if testHelper.Received != expected {
			t.Errorf("Expected result was incorrect. %s != %s",
				testHelper.Received, expected)
		}
	})
}