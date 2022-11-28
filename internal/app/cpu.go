package app

import (
	"fmt"
	"time"
)

type CPU struct {
	a       Register
	b       Register
	carry   Register
	pc      Register
	decoder Decoder
	rom     ROM
	in      IO
	out     IO
}

func NewCPU() CPU {
	return CPU{
		a:       NewRegister(),
		b:       NewRegister(),
		carry:   NewRegister(),
		pc:      NewRegister(),
		decoder: NewDecoder(),
		rom:     NewROM(),
		in:      NewIO(),
		out:     NewIO(),
	}
}

func (c *CPU) fetch() string {
	fmt.Println("Fetch")
	return c.rom.Fetch()
}

func (c *CPU) decode(instruction string) (Opecode, Immidiate) {
	fmt.Println("Decode")
	return c.decoder.Decode(instruction)
}

func (c *CPU) execute(opecode Opecode, immidiate Immidiate) {
	fmt.Println("Execute")

	// TODO main processing
	switch opecode {
	case MOV_A:
		// TODO MOV_A function call
	case MOV_B:
		// TODO ...
	default:
		fmt.Println("Error opecode exist!")
	}
}

func (c *CPU) Run() {
	for {
		instruction := c.fetch()
		opecode, immidiate := c.decode(instruction)
		c.execute(opecode, immidiate)

		time.Sleep(time.Millisecond * 100)
	}
}
