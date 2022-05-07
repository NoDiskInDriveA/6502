package cpu

import "time"

type Clock interface {
	Ticker() chan bool
}

type TimedClock struct {
	tick chan bool
}

type StepClock struct {
	tick chan bool
}

func NewTimedClock() Clock {
	tick := make(chan bool)
	ticker := time.NewTicker(1 * time.Millisecond)
	go func() {
		for range ticker.C {
			tick <- true
		}
	}()

	return &TimedClock{tick}
}

func (tc *TimedClock) Ticker() chan bool {
	return tc.tick
}

func NewStepClock() Clock {
	tick := make(chan bool)
	go func() {

	}()

	return &TimedClock{tick}
}

func (sc *StepClock) Ticker() chan bool {
	return sc.tick
}
