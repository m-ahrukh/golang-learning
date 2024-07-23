package mocking

import (
	"fmt"
	"io"
	"os"
	"time"
)

// func Countdown(out io.Writer) {
// 	fmt.Fprint(out, "3")
// }

// func Countdown(out io.Writer) {
// 	for i := countdownStart; i > 0; i-- {
// 		fmt.Fprintln(out, i)
// 	}
// 	fmt.Fprint(out, finalWord)
// }

type Sleeper interface {
	Sleep()
}

// type DefaultSleeper struct{}

// func (d *DefaultSleeper) Sleep() {
// 	time.Sleep(1 * time.Second)
// }

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

const (
	finalWord      = "Go!"
	countdownStart = 3
)

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}

func main() {
	// Countdown(os.Stdout)
	// sleeper := &DefaultSleeper{}

	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
