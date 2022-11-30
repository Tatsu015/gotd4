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
	input   IO
	output  IO
}

func NewCPU(rom ROM, input IO, output IO) CPU {
	return CPU{
		a:       NewRegister(),
		b:       NewRegister(),
		carry:   NewRegister(),
		pc:      NewRegister(),
		decoder: NewDecoder(),
		rom:     rom,
		input:   input,
		output:  output,
	}
}

func (c *CPU) add_a(i Immidiate)  { /*fmt.Println(i)*/ }
func (c *CPU) mov_ab(i Immidiate) { /*fmt.Println(i)*/ }
func (c *CPU) in_a(i Immidiate)   { /*fmt.Println(i)*/ }
func (c *CPU) mov_a(i Immidiate)  { /*fmt.Println(i)*/ }
func (c *CPU) mov_ba(i Immidiate) { /*fmt.Println(i)*/ }
func (c *CPU) add_b(i Immidiate)  { /*fmt.Println(i)*/ }
func (c *CPU) in_b(i Immidiate)   { /*fmt.Println(i)*/ }
func (c *CPU) mov_b(i Immidiate)  { /*fmt.Println(i)*/ }
func (c *CPU) out_b(i Immidiate)  { /*fmt.Println(i)*/ }
func (c *CPU) out(i Immidiate)    { /*fmt.Println(i)*/ }
func (c *CPU) jmp(i Immidiate)    { /*fmt.Println(i)*/ }
func (c *CPU) jnc(i Immidiate)    { /*fmt.Println(i)*/ }

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
	ope, imm := c.decoder.Decode(i)
	fmt.Printf("[Decode] Ins: %08b Ope: %04b Imm: %04b\n", i, ope, imm)
	return ope, imm
}

func (c *CPU) execute(o Opecode, i Immidiate) error {
	switch o {
	case ADD_A:
		c.add_a(i)
		return nil
	case MOV_AB:
		c.mov_ab(i)
		return nil
	case IN_A:
		c.in_a(i)
		return nil
	case MOV_A:
		c.mov_a(i)
		return nil
	case MOV_BA:
		c.mov_ba(i)
		return nil
	case ADD_B:
		c.add_b(i)
		return nil
	case IN_B:
		c.in_b(i)
		return nil
	case MOV_B:
		c.mov_b(i)
		return nil
	case OUT_B:
		c.out_b(i)
		return nil
	case OUT:
		c.out(i)
		return nil
	case JMP:
		c.jmp(i)
		return nil
	case JNC:
		c.jnc(i)
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
