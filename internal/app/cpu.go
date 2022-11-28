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

func (c *CPU) fetch() {
	fmt.Println("Fetch")
}

func (c *CPU) decode() {
	fmt.Println("Decode")
}

func (c *CPU) execute() {
	fmt.Println("Execute")
}

func (c *CPU) Run() {
	for {
		c.fetch()
		c.decode()
		c.execute()
		time.Sleep(time.Millisecond * 100)
	}
}
