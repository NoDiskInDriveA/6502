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
	StartClock(tickDuration time.Duration) chan bool
}

type clockAwareDeviceDefinition struct {
	device  ClockAwareDevice
	divider int
}

type synchronizedSystemClock struct {
	devices  []clockAwareDeviceDefinition
	debugger ClockAwareDevice
}

func (tc *synchronizedSystemClock) AttachDevice(device ClockAwareDevice, divider int) {
	if divider != 1 {
		panic("Clock dividers other than 1 are currently not implemented.")
	}
	tc.devices = append(tc.devices, clockAwareDeviceDefinition{device, divider})
}

// debugger which will run _after_ each cycle
func (tc *synchronizedSystemClock) AttachDebugger(device ClockAwareDevice) {
	tc.debugger = device
}

func (tc *synchronizedSystemClock) StartClock(tickDuration time.Duration) chan bool {
	ticker := time.NewTicker(tickDuration)
	stop := make(chan bool)
	running := true
	deviceCount := len(tc.devices)
	var waitgroup sync.WaitGroup
	go func() {
		for running {
			<-ticker.C
			if tc.debugger != nil {
				running = running && tc.debugger.Tick()
			}
			waitgroup.Add(deviceCount)
			for _, d := range tc.devices {
				go func(device ClockAwareDevice) {
					running = running && device.Tick()
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

func NewSynchronizedClock() SystemClock {
	return &synchronizedSystemClock{}
}
