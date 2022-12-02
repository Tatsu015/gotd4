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

func (c *CPU) showOutput() {
	v := c.output.value()
	fmt.Printf("%04b\n", v)
}

func (c *CPU) add_a(i Immidiate) {
	oldVal := c.a.value()
	newVal := oldVal + i
	if newVal > REGISTER_CAPACITY {
		uncarried := 0x0f & newVal
		c.a.setValue(uncarried)
		c.carry.setValue(1)
	} else {
		c.a.setValue(newVal)
		c.carry.setValue(0)
	}
	c.a.setValue(newVal)
}

func (c *CPU) mov_ab() {
	im := c.b.value()
	c.a.setValue(im)
	c.carry.setValue(0)

}

func (c *CPU) in_a() {
	i := c.input.value()
	c.a.setValue(i)
	c.carry.setValue(0)
}

func (c *CPU) mov_a(i Immidiate) {
	c.a.setValue(i)
	c.carry.setValue(0)
}

func (c *CPU) mov_ba() {
	im := c.a.value()
	c.b.setValue(im)
	c.carry.setValue(0)
}

func (c *CPU) add_b(i Immidiate) {
	oldVal := c.b.value()
	newVal := oldVal + i
	if newVal > REGISTER_CAPACITY {
		uncarried := 0x0f & newVal
		c.b.setValue(uncarried)
		c.carry.setValue(1)
	} else {
		c.b.setValue(newVal)
		c.carry.setValue(0)
	}
	c.b.setValue(newVal)
}

func (c *CPU) in_b() {
	i := c.input.value()
	c.b.setValue(i)
	c.carry.setValue(0)
}

func (c *CPU) mov_b(i Immidiate) {
	c.b.setValue(i)
	c.carry.setValue(0)
}

func (c *CPU) out_b() {
	i := c.b.value()
	c.output.setValue(i)
	c.carry.setValue(0)
}

func (c *CPU) out(i Immidiate) {
	c.output.setValue(i)
	c.carry.setValue(0)
}

func (c *CPU) jmp(i Immidiate) {
	c.pc.setValue(i)
	c.carry.setValue(0)
}

func (c *CPU) jnc(i Immidiate) {
	if c.carry.value() == 0 {
		c.pc.setValue(i)
	}
	c.carry.setValue(0)
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
	// fmt.Printf("[Fetch] %v\n", ins)
	return ins
}

func (c *CPU) decode(i Instruction) (Opecode, Immidiate) {
	ope, imm := c.decoder.Decode(i)
	// fmt.Printf("[Decode] Ins: %08b Ope: %04b Imm: %04b\n", i, ope, imm)
	return ope, imm
}

func (c *CPU) execute(o Opecode, i Immidiate) error {
	switch o {
	case ADD_A:
		c.add_a(i)
		return nil
	case MOV_AB:
		c.mov_ab()
		return nil
	case IN_A:
		c.in_a()
		return nil
	case MOV_A:
		c.mov_a(i)
		return nil
	case MOV_BA:
		c.mov_ba()
		return nil
	case ADD_B:
		c.add_b(i)
		return nil
	case IN_B:
		c.in_b()
		return nil
	case MOV_B:
		c.mov_b(i)
		return nil
	case OUT_B:
		c.out_b()
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
		return fmt.Errorf("opecode %v not exist!", o)
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
		c.showOutput()
		c.waitClockUp()
		c.incrementPC()
	}
}
