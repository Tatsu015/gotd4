package emulator

import (
	"fmt"
	"time"
)

type Emulator struct {
	cpu    CPU
	rom    *ROM
	input  *Input
	output *Output
}

func NewEmulator(rom *ROM, input *Input, output *Output) Emulator {
	cpu := NewCPU(rom, input, output)
	return Emulator{cpu, rom, input, output}
}

func (e *Emulator) showOutput() {
	v := e.output.value()
	fmt.Printf("%04b\n", v)
}

func (c *Emulator) waitClockUp() {
	time.Sleep(time.Millisecond * 100)
}

func (e *Emulator) Run() {
	for {
		e.cpu.Progress()

		// wait and PC count up for next loop
		e.showOutput()
		e.waitClockUp()
	}
}
