package main

import (
	"io"
	"fmt"
	"os"
	"time"
)

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}

const finalWord = "GO!"
const countdownStart = 3

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}	

	fmt.Fprint(out, finalWord)
}

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct {}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}