package cpu

import "time"

type Clock interface {
	Ticker() chan bool
}

type TimedClock struct {
	tick chan bool
}

func NewTimedClock() Clock {
	tick := make(chan bool)
	ticker := time.NewTicker(10 * time.Millisecond)
	go func() {
		for {
			select {
			case <-ticker.C:
				tick <- true
			}
		}
	}()

	return &TimedClock{tick}
}

func (tc *TimedClock) Ticker() chan bool {
	return tc.tick
}
