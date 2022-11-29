package app

import (
	"fmt"
	"time"
)

type Opecode int16

const (
	ADD_A  Opecode = 0x00
	MOV_AB Opecode = 0x01
	IN_A   Opecode = 0x02
	MOV_A  Opecode = 0x03
	MOV_BA Opecode = 0x04
	ADD_B  Opecode = 0x05
	IN_B   Opecode = 0x06
	MOV_B  Opecode = 0x07
	OUT_B  Opecode = 0x09
	OUT    Opecode = 0x0b
	JMP    Opecode = 0x0e
	JNC    Opecode = 0x0f
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

func NewCPU(rom ROM, in IO, out IO) CPU {
	return CPU{
		a:       NewRegister(),
		b:       NewRegister(),
		carry:   NewRegister(),
		pc:      NewRegister(),
		decoder: NewDecoder(),
		rom:     rom,
		in:      in,
		out:     out,
	}
}

func (c *CPU) waitClockUp() {
	time.Sleep(time.Millisecond * 100)
}

func (c *CPU) incrementPC() {
	v := c.pc.value()
	c.pc.setValue(v + 1)
}

func (c *CPU) fetch(a Adress) Instruction {
	ins := c.rom.Fetch(a)
	fmt.Printf("[Fetch] %v\n", ins)
	return ins
}

func (c *CPU) decode(i Instruction) (Opecode, Immidiate) {
	fmt.Println("Decode")
	return c.decoder.Decode(i)
}

func (c *CPU) execute(o Opecode, i Immidiate) error {
	fmt.Println("Execute")

	// TODO main processing
	switch o {
	case MOV_A:
		// TODO MOV_A function call
		return nil
	case MOV_B:
		// TODO ...
		return nil
	default:
		return nil // TODO for motion check
		// return fmt.Errorf("opecode %v not exist!", o)
	}
}

func (c *CPU) Run() {
	for {
		// fetch program from ROM
		ad := Adress(c.pc.value())
		inst := c.fetch(ad)

		// analyze fetch data
		ope, imm := c.decode(inst)

		// execute program
		err := c.execute(ope, imm)
		if err != nil {
			fmt.Println(err)
			return
		}

		// wait and PC count up for next loop
		c.waitClockUp()
		c.incrementPC()
	}
}
