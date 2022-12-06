package emulator

import (
	"time"

	"github.com/Tatsu015/gotd4.git/internal/emulator/cpu"
	"github.com/Tatsu015/gotd4.git/internal/emulator/debugger"
	"github.com/Tatsu015/gotd4.git/internal/emulator/io"
	"github.com/Tatsu015/gotd4.git/internal/emulator/rom"
)

type Emulator struct {
	cpu      *cpu.CPU
	rom      *rom.ROM
	input    *io.Input
	output   *io.Output
	clock    time.Duration
	debugger *debugger.Debugger
}

func NewEmulator(rom *rom.ROM, input *io.Input, output *io.Output, clock time.Duration, useDebugger bool) Emulator {
	cpu := cpu.NewCPU(rom, input, output)

	if useDebugger {
		d := debugger.NewDebugger(&cpu)
		return Emulator{&cpu, rom, input, output, clock, &d}
	} else {
		return Emulator{&cpu, rom, input, output, clock, nil}
	}
}

func (c *Emulator) waitClockUp() {
	time.Sleep(time.Millisecond * c.clock)
}

func (e *Emulator) Run() {
	for {
		e.cpu.Progress()
		if e.debugger != nil {
			e.debugger.Break()
		}
		// wait and PC count up for next loop
		e.output.Show()
		e.waitClockUp()
	}
}
