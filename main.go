package main

import (
	"go_tdd/Mocking"
	"os"
)

func main() {
	sleeper := &Mocking.DefaultSleeper{}
	Mocking.CountDown(os.Stdout, sleeper)
}
