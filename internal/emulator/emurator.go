package emulator

import (
	"time"

	"github.com/Tatsu015/gotd4.git/internal/emulator/cpu"
	"github.com/Tatsu015/gotd4.git/internal/emulator/io"
	"github.com/Tatsu015/gotd4.git/internal/emulator/rom"
)

type Emulator struct {
	cpu    cpu.CPU
	rom    *rom.ROM
	input  *io.Input
	output *io.Output
	clock  time.Duration
}

func NewEmulator(rom *rom.ROM, input *io.Input, output *io.Output, clock time.Duration) Emulator {
	cpu := cpu.NewCPU(rom, input, output)
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
