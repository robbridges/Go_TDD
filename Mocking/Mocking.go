package Mocking

import (
	"fmt"
	"io"
	"time"
)

var countDownStart = []string{"3", "2", "1"}

const finalWord = "Go!"

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type ConfigurableSleeper struct {
	Duration   time.Duration
	TimedSleep func(duration time.Duration)
}

type Sleeper interface {
	Sleep()
}

func CountDown(
	w io.Writer,
	sleeper Sleeper,
) {
	for _, v := range countDownStart {
		fmt.Fprintln(w, v)
		sleeper.Sleep()
	}

	fmt.Fprint(w, finalWord)
}

func (c *ConfigurableSleeper) Sleep() {
	c.TimedSleep(c.Duration)
}
