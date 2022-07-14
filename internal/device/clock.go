package device

import (
	"sync"
	"time"
)

type ClockAwareDevice interface {
	Tick() bool
}

type SystemClock interface {
	AttachDevice(device ClockAwareDevice, divider int)
	AttachDebugger(device ClockAwareDevice)
	RunClock(tickDuration time.Duration) chan bool
	TickClock()
}

type clockAwareDeviceDefinition struct {
	device  ClockAwareDevice
	divider int
}

type synchronizedSystemClock struct {
	devices     []clockAwareDeviceDefinition
	debugger    ClockAwareDevice
	deviceCount int
	running     bool
}

func (tc *synchronizedSystemClock) AttachDevice(device ClockAwareDevice, divider int) {
	if divider != 1 {
		panic("Clock dividers other than 1 are currently not implemented.")
	}
	tc.devices = append(tc.devices, clockAwareDeviceDefinition{device, divider})
	tc.deviceCount = len(tc.devices)
}

// debugger which will run _after_ each cycle
func (tc *synchronizedSystemClock) AttachDebugger(device ClockAwareDevice) {
	tc.debugger = device
}

func (tc *synchronizedSystemClock) RunClock(tickDuration time.Duration) chan bool {
	ticker := time.NewTicker(tickDuration)
	stop := make(chan bool)
	var waitgroup sync.WaitGroup
	go func() {
		for tc.running {
			<-ticker.C
			if tc.debugger != nil {
				tc.running = tc.running && tc.debugger.Tick()
			}
			waitgroup.Add(tc.deviceCount)
			for _, d := range tc.devices {
				go func(device ClockAwareDevice) {
					tc.running = tc.running && device.Tick()
					waitgroup.Done()
				}(d.device)
			}
			waitgroup.Wait()
		}
		ticker.Stop()
		stop <- true
	}()

	return stop
}

func (tc *synchronizedSystemClock) TickClock() {
	var waitgroup sync.WaitGroup
	waitgroup.Add(tc.deviceCount)
	for _, d := range tc.devices {
		go func(device ClockAwareDevice) {
			tc.running = tc.running && device.Tick()
			waitgroup.Done()
		}(d.device)
	}
	waitgroup.Wait()
}

func NewSynchronizedClock() SystemClock {
	return &synchronizedSystemClock{running: true}
}
