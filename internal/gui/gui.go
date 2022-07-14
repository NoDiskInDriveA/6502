package gui

import (
	"fmt"

	"github.com/NoDiskInDriveA/6502/internal/processor/mos_6502"
	"github.com/jroimartin/gocui"
)

type MemoryView struct {
	monitor *mos_6502.Monitor
}

func NewMemoryView(monitor *mos_6502.Monitor) *MemoryView {
	return &MemoryView{monitor}
}

func (mv *MemoryView) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	v, err := g.SetView("memory", 0, 0, maxX-20, maxY-1)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v.Clear()
	fmt.Fprint(v, mv.monitor.GetMemoryView())

	return nil
}

type StatusView struct {
	monitor *mos_6502.Monitor
}

func NewStatusView(monitor *mos_6502.Monitor) *StatusView {
	return &StatusView{monitor}
}

func (sv *StatusView) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	v, err := g.SetView("status", maxX-19, 3, maxX-1, maxY-1)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v.Clear()
	for _, s := range sv.monitor.GetInternalStatus() {
		fmt.Fprintln(v, s)
	}

	return nil
}

type StatusWordView struct {
	monitor *mos_6502.Monitor
}

func NewStatusWordView(monitor *mos_6502.Monitor) *StatusWordView {
	return &StatusWordView{monitor}
}

func (swv *StatusWordView) Layout(g *gocui.Gui) error {
	maxX, _ := g.Size()

	v, err := g.SetView("status_word", maxX-19, 0, maxX-1, 2)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v.Clear()
	fmt.Fprint(v, swv.monitor.GetStatusWord())

	return nil
}
