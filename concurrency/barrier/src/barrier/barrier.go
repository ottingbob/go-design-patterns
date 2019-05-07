package barrier

// Pattern puts up a barrier so execution may not pass
// until we have all the results we need.

// Controls the value of a type with the data coming from
// one or more Goroutines

// Control the correctness of any of those incoming data
// pipes so that no inconsistent data is returned. We don't
// want a partially filled result because one of the pipes
// has returned an error

import (
	"bytes"
	"io"
	"os"

	"io/ioutil"
	"net/http"
	"fmt"
	"time"
)

func CaptureBarrierOutput(endpoints ...string) string {
	reader, writer, _ := os.Pipe()

	os.Stdout = writer

	out := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, reader)
		out <- buf.String()
	}()

	barrier(endpoints...)

	writer.Close()
	temp := <-out

	return temp
}

type barrierResp struct {
	Err error
	Resp string
}

var TimeoutMilliseconds int = 5000

func barrier(endpoints ...string) {
	requestNumber := len(endpoints)

	in := make(chan barrierResp, requestNumber)
	defer close(in)

	responses := make([]barrierResp, requestNumber)

	for _, endpoint := range endpoints {
		go makeRequest(in, endpoint)
	}

	var hasError bool
	for i := 0; i < requestNumber; i++ {
		resp := <-in
		if resp.Err != nil {
			fmt.Println("ERROR: ", resp.Err)
			hasError = true
			break
		}
		responses[i] = resp
	}

	if !hasError {
		for _, resp := range responses {
			fmt.Println(resp.Resp)
		}
	}
}

func makeRequest(out chan<- barrierResp, url string) {
	res := barrierResp{}
	client := http.Client {
		Timeout: time.Duration(time.Duration(TimeoutMilliseconds) * time.Millisecond),
	}

	resp, err := client.Get(url)
	if err != nil {
		res.Err = err
		out <- res
		return
	}

	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		res.Err = err
		out <- res
		return
	}

	res.Resp = string(byt)
	out <- res
}