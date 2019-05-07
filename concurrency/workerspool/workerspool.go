package main

import (
	"fmt"
	"sync"
	"strings"
	"log"
	"time"
)

type Request struct {
	Data interface{}
	Handler RequestHandler
}

type RequestHandler func(interface{})

func NewStringRequest(s string, id int, wg *sync.WaitGroup) Request {
	myRequest := Request{
		// Data: "Hello", Handler: func(i interface{}) {
		Data: s, Handler: func(i interface{}) {
			defer wg.Done()
			s, ok := i.(string)
			if !ok {
				log.Fatal("Invalid casting to string")
			}
			fmt.Println(s)
		},
	}

	return myRequest
}

type WorkerLauncher interface {
	LaunchWorker(id int, in chan Request)
}

type Dispatcher interface {
	LaunchWorker(id int, w WorkerLauncher)
	MakeRequest(Request)
	Stop()
}

type dispatcher struct {
	inCh chan Request
}

func (d *dispatcher) LaunchWorker(id int, w WorkerLauncher) {
	w.LaunchWorker(id, d.inCh)
}

func (d *dispatcher) Stop() {
	close(d.inCh)
}

func (d *dispatcher) MakeRequest(r Request) {
	select {
	case d.inCh <- r:
	case <-time.After(time.Second * 5):
		return
	}
}

func NewDispatcher(b int) Dispatcher {
	return &dispatcher{
		inCh: make(chan Request, b),
	}
}

type PrefixSuffixWorker struct {
	id int
	prefixS string
	suffixS string
}

func (w *PrefixSuffixWorker) LaunchWorker(i int, in chan Request) {
	w.prefix(w.append(w.uppercase(in)))
}

func (w *PrefixSuffixWorker) uppercase(in <-chan Request) <-chan Request {
	out := make(chan Request)
	
	go func() {
		for msg := range in {
			s, ok := msg.Data.(string)

			if !ok {
				msg.Handler(nil)
				continue
			}

			msg.Data = strings.ToUpper(s)

			out <- msg
		}

		close(out)
	}()

	return out
}

func (w *PrefixSuffixWorker) append(in <-chan Request) <-chan Request {
	out := make(chan Request)
	go func() {
		for msg := range in {
			uppercaseString, ok := msg.Data.(string)
			if !ok {
				msg.Handler(nil)
				continue
			}
			msg.Data = fmt.Sprintf("%s%s", uppercaseString, w.suffixS)
			out <- msg
		}
		close(out)
	}()
	return out
}

func (w *PrefixSuffixWorker) prefix(in <-chan Request) {
	go func() {
		for msg := range in {
			uppercasedStringWithSuffix, ok := msg.Data.(string)

			if !ok {
				msg.Handler(nil)
				continue
			}

			msg.Handler(fmt.Sprintf("%s%s", w.prefixS, 
				uppercasedStringWithSuffix))
		}
	}()
}

func main() {
	bufferSize := 100
	var dispatcher Dispatcher = NewDispatcher(bufferSize)

	workers := 3
	for i := 0; i < workers; i++ {
		var w WorkerLauncher = &PrefixSuffixWorker{
			prefixS: fmt.Sprintf("WorkerID: %d -> ", i),
			suffixS: " World",
			id: i,
		}
		dispatcher.LaunchWorker(i, w)
	}

	requests := 10
	var wg sync.WaitGroup
	wg.Add(requests)

	for i := 0; i < requests; i++ {
		req := NewStringRequest(fmt.Sprintf("(Msg_id: %d) -> Hello", i), i, &wg)
		dispatcher.MakeRequest(req)
	}

	dispatcher.Stop()
	wg.Wait()
}