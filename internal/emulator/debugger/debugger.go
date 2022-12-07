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

func (d *Debugger) PrintCPUStatus() {
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	d.cpu.Show()
}

func (d *Debugger) Break() {
	fmt.Println("c + Enter: Go to next line")
	fmt.Scan(&d.currentCmd)
}
