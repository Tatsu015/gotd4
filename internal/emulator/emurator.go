package emulator

import (
	"time"

	"github.com/Tatsu015/gotd4.git/internal/emulator/cpu"
	"github.com/Tatsu015/gotd4.git/internal/emulator/debugger"
	"github.com/Tatsu015/gotd4.git/internal/emulator/io"
	"github.com/Tatsu015/gotd4.git/internal/emulator/rom"
)

type Emulator struct {
	cpu         *cpu.CPU
	rom         *rom.ROM
	input       *io.Input
	output      *io.Output
	clock       time.Duration
	debugger    *debugger.Debugger
	useDebugger bool
	verbose     bool
}

func NewEmulator(rom *rom.ROM, input *io.Input, output *io.Output, clock time.Duration, useDebugger bool, verbose bool) Emulator {
	cpu := cpu.NewCPU(rom, input, output)
	d := debugger.NewDebugger(&cpu)
	return Emulator{&cpu, rom, input, output, clock, &d, useDebugger, verbose}
}

func (c *Emulator) waitClockUp() {
	time.Sleep(time.Millisecond * c.clock)
}

func (e *Emulator) Run() {
	// clear console display first
	// e.output.InitDisplay()

	for {
		e.cpu.Progress()

		if !e.useDebugger {
			e.output.Clear()
		}
		if e.verbose {
			e.debugger.PrintCPUStatus()
		}
		if e.useDebugger {
			e.debugger.Break()
		}
		e.output.Show()
		// wait and PC count up for next loop
		e.waitClockUp()
	}
}
