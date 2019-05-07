package adapter

import (
	"testing"

	a "adapter"
)

func TestAdapter(t *testing.T) {
	msg := "Hello World!"
	adapter := a.PrinterAdapter{
		OldPrinter: &a.MyLegacyPrinter{},
		Msg: msg, 
	}

	returnedMsg := adapter.PrintStored()
	if returnedMsg != "Legacy Printer: Adapter: Hello World!\n" {
		t.Errorf("Message didn't match: %s\n", returnedMsg)
	}

	adapter = a.PrinterAdapter{
		OldPrinter: nil,
		Msg: msg,
	}
	
	returnedMsg = adapter.PrintStored()
	if returnedMsg != "Hello World!" {
		t.Errorf("Message didn't match: %s\n", returnedMsg)
	}
}