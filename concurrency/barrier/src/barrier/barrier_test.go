package barrier

import (
	"strings"
	"testing"

	b "barrier"
)

func TestBarrier(t *testing.T) {
	t.Run("Correct endpoints", func(t *testing.T) {
		endpoints := []string{"http://httpbin.org/headers", "http://httpbin.org/User-Agent"}

		result := b.CaptureBarrierOutput(endpoints...)

		if !strings.Contains(result, "Accept-Encoding") || !strings.Contains(result, "User-Agent") {
			t.Fail()
		}

		t.Log(result)
	})

	t.Run("One endpoint incorrect", func(t *testing.T) {
		endpoints := []string{"http://malformed-url", "http://httpbin.org/User-Agent"}

		result := b.CaptureBarrierOutput(endpoints...)

		if !strings.Contains(result, "ERROR") {
			t.Fail()
		}

		t.Log(result)
	})

	t.Run("Very short timeout", func(t *testing.T) {
		endpoints := []string{"http://httpbin.org/headers", "http://httpbin.org/User-Agent"}
		b.TimeoutMilliseconds = 1

		result := b.CaptureBarrierOutput(endpoints...)

		if !strings.Contains(result, "Timeout") {
			t.Fail()
		}

		t.Log(result)
	})
}