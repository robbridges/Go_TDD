package main

import (
	"go_tdd/Mocking"
	"os"
	"time"
)

func main() {
	sleeper := &Mocking.ConfigurableSleeper{5 * time.Second, time.Sleep}
	Mocking.CountDown(os.Stdout, sleeper)
}
