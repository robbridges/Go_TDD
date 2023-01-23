package Mocking

import (
	"bytes"
	"testing"
)

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountDown(t *testing.T) {
	buffer := &bytes.Buffer{}
	sleeper := &SpySleeper{}
	CountDown(buffer, sleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}

	if sleeper.Calls != 3 {
		t.Errorf("The sleep function was not called 3 times")
	}
}
