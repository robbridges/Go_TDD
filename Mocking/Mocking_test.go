package Mocking

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

const write = "write"
const sleep = "sleep"

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(duration time.Duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

type SpySleeper struct {
	Calls int
}

type SpyCountDownOperations struct {
	Calls []string
}

func TestCountDown(t *testing.T) {
	t.Run("Happy path spy sleeper", func(t *testing.T) {
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
	})

	t.Run("Should sleep before printing", func(t *testing.T) {
		spySleepPrinter := &SpyCountDownOperations{}

		CountDown(spySleepPrinter, spySleepPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("Wanted calls to be %v, got %v", want, spySleepPrinter.Calls)
		}
	})

	t.Run("Test Configruable sleeper", func(t *testing.T) {
		sleepTime := 5 * time.Second

		spyTime := &SpyTime{}
		sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
		sleeper.Sleep()

		if spyTime.durationSlept != sleepTime {
			t.Errorf("The sleeper did not sleep for 5 seconds as expected")
		}
	})
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func (s *SpyCountDownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

// we need to give it a write  method that extends io.Writer so that it can be passed as an io writer
func (s *SpyCountDownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}
