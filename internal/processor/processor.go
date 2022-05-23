package processor

import "github.com/NoDiskInDriveA/6502/internal/device"

type Processor interface {
	device.ClockAwareDevice
	Reset()
}
