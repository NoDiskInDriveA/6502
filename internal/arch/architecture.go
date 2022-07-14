package arch

import "github.com/NoDiskInDriveA/6502/internal/processor/mos_6502"

type Architecture interface {
	Run()
	Step()
}

type MonitoredArchitecture interface {
	Architecture
	GetMonitor() *mos_6502.Monitor
}
