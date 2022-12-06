package debugger

import (
	"fmt"

	"github.com/Tatsu015/gotd4.git/internal/emulator/cpu"
)

type Debugger struct {
	cpu        *cpu.CPU
	currentCmd string
}

func NewDebugger(cpu *cpu.CPU) Debugger {
	return Debugger{
		cpu: cpu,
	}
}

func (d *Debugger) Break() {
	fmt.Println("-----------------------------------------------------")
	d.cpu.Show()
	fmt.Println("c + Enter: Go to next line")
	fmt.Scan(&d.currentCmd)
}
