package main

import (
	"io"
	"fmt"
	"os"
	"time"
)

func main() {
	Countdown(os.Stdout)
}

const finalWord = "GO!"
const countdownStart = 3

func Countdown(out io.Writer) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		time.Sleep(1 * time.Second)
	}	

	fmt.Fprint(out, finalWord)
}
