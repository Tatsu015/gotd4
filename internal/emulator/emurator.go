package emulator

import (
	"time"
)

type Emulator struct {
	cpu    CPU
	rom    *ROM
	input  *Input
	output *Output
	clock  time.Duration
}

func NewEmulator(rom *ROM, input *Input, output *Output, clock time.Duration) Emulator {
	cpu := NewCPU(rom, input, output)
	return Emulator{cpu, rom, input, output, clock}
}

func (c *Emulator) waitClockUp() {
	time.Sleep(time.Millisecond * c.clock)
}

func (e *Emulator) Run() {
	for {
		e.cpu.Progress()

		// wait and PC count up for next loop
		e.output.Show()
		e.waitClockUp()
	}
}
