package bridge

import (
	"testing"
	"strings"

	b "bridge"
)

func TestPrintAPI1(t *testing.T) {
	api1 := b.PrinterImpl1{}

	err := api1.PrintMessage("Hello")
	if err != nil {
		t.Errorf("Error trying to use the API1 implementation: Message: %s\n",
			err.Error())
	}
}

func TestPrintAPI2(t *testing.T) {
	api2 := b.PrinterImpl2{}

	err := api2.PrintMessage("Hello")
	if err != nil {
		expectedErrorMsg := "You need to pass an io.Writer to PrinterImpl2"
		if !strings.Contains(err.Error(), expectedErrorMsg) {
			t.Errorf("Error message was not correct.\n" +
				"Actual: %s\nExpected: %s\n", err.Error(), expectedErrorMsg)
		}
	}

	testWriter := b.TestWriter{}
	api2 = b.PrinterImpl2{
		Writer: &testWriter,
	}

	expectedMessage := "Hello"
	err = api2.PrintMessage(expectedMessage)
	if err != nil {
		t.Errorf("Error trying to use the API2 implementation: %s\n", err.Error())
	}

	if testWriter.Msg != expectedMessage {
		t.Fatalf("API2 did not write correctly on the io.Writer.\n" +
			"Actual: %s\nExpected: %s\n", testWriter.Msg, expectedMessage)
	}
}

func TestNormalPrinter_Print(t *testing.T) {
	expectedMessage := "Hello io.Writer"

	normal := b.NormalPrinter{
		Msg: expectedMessage,
		Printer: &b.PrinterImpl1{},
	}

	err := normal.Print()
	if err != nil {
		t.Errorf(err.Error())
	}

	testWriter := b.TestWriter{}
	normal = b.NormalPrinter{
		Msg: expectedMessage,
		Printer: &b.PrinterImpl2{
			Writer: &testWriter,
		},
	}

	err = normal.Print()
	if err != nil {
		t.Error(err.Error())
	}

	if testWriter.Msg != expectedMessage {
		t.Errorf("The expected message on the io.Writer doesn't match actual.\n" + 
			"Actual: %s\nExpected: %s\n", testWriter.Msg, expectedMessage)
	}
}

func TestPacktPrinter_Print(t *testing.T) {
	passedMessage := "Hello io.Writer"
	expectedMessage := "Message from Packt: Hello io.Writer"

	packt := b.PacktPrinter{
		Msg: passedMessage,
		Printer: &b.PrinterImpl1{},
	}

	err := packt.Print()
	if err != nil {
		t.Errorf(err.Error())
	}

	testWriter := b.TestWriter{}
	packt = b.PacktPrinter{
		Msg: passedMessage,
		Printer: &b.PrinterImpl2{
			Writer: &testWriter,
		},
	}

	err = packt.Print()
	if err != nil {
		t.Error(err.Error())
	}

	if testWriter.Msg != expectedMessage {
		t.Errorf("The expected message on the io.Writer doesn't match actual.\n" + 
			"Actual: %s\nExpected: %s\n", testWriter.Msg, expectedMessage)
	}
}