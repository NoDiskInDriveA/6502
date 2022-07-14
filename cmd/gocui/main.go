package main

import (
	"log"

	"github.com/NoDiskInDriveA/6502/internal/arch"
	"github.com/NoDiskInDriveA/6502/internal/gui"
	"github.com/jroimartin/gocui"
)

func main() {
	architecture := arch.NewMonitoredFantasyArchitecture()

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManager(
		gui.NewMemoryView(architecture.GetMonitor()),
		gui.NewStatusView(architecture.GetMonitor()),
		gui.NewStatusWordView(architecture.GetMonitor()),
	)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", 's', gocui.ModNone, step(architecture)); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", 'r', gocui.ModNone, run(architecture)); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func step(a arch.Architecture) func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		a.Step()
		return nil
	}
}

func run(a arch.Architecture) func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		a.Run()
		return nil
	}
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
