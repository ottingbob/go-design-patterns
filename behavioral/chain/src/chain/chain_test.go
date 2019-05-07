package chain

import (
	"fmt"
	"strings"
	"testing"

	c "chain"
)

type myTestWriter struct {
	receivedMessage *string
}

func (m *myTestWriter) Write(p []byte) (int, error) {
	tempMessage := string(p)
	m.receivedMessage = &tempMessage
	return len(p), nil
}

func TestCreateDefaultChain(t *testing.T) {
	myWriter := myTestWriter{}
	writerLogger := c.WriterLogger{Writer: &myWriter}
	second := c.SecondLogger{NextChain: &writerLogger}
	chain := c.FirstLogger{NextChain: &second}

	t.Run("3 loggers, 2 of them writes to console, second only if it finds " +
		"the world 'hello', third writes to some variabel if second found 'hello'",
		func(t *testing.T) {
			chain.Next("message that breaks the chain")

			if myWriter.receivedMessage != nil {
				t.Error("Last link should not receive any message")
			}

			chain.Next("Hello")

			if myWriter.receivedMessage == nil || !strings.Contains(*myWriter.receivedMessage, "Hello") {
				t.Fatal("Last link didn't receive expected message")
			}
		})
	
	t.Run("2 loggers, second uses the closure implementation", func(t *testing.T) {
		myWriter = myTestWriter{}
		closureLogger := c.ClosureChain{
			Closure: func(s string) {
				fmt.Printf("My closure logger! Message %s\n", s)
				myWriter.receivedMessage = &s
			},
		}

		writerLogger.NextChain = &closureLogger

		chain.Next("Hello closure logger")
		
		if *myWriter.receivedMessage != "Hello closure logger" {
			t.Fatal("Expected message wasn't received in myWriter")
		}
	})
}