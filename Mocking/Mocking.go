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

func CountDown(w io.Writer, sleeper Sleeper) {
	for _, v := range countDownStart {
		fmt.Fprintln(w, v)
		time.Sleep(1 * time.Second)
	}
	fmt.Fprint(w, finalWord)
}
