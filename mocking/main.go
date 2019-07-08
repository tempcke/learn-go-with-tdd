package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		delayedOutput(strconv.Itoa(i), writer, sleeper)
	}
	delayedOutput(finalWord, writer, sleeper)
}

func delayedOutput(output string, writer io.Writer, sleeper Sleeper) {
	sleeper.Sleep()
	_, _ = fmt.Fprintln(writer, output)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
