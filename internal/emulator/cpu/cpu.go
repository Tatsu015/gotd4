package cpu

import (
	"fmt"

	"github.com/Tatsu015/gotd4.git/internal/define"
	"github.com/Tatsu015/gotd4.git/internal/emulator/io"
	"github.com/Tatsu015/gotd4.git/internal/emulator/rom"
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
	rom     *rom.ROM
	input   *io.Input
	output  *io.Output
}

func NewCPU(rom *rom.ROM, input *io.Input, output *io.Output) CPU {
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

func (c *CPU) add_a(i define.Immidiate) {
	oldVal := c.a.value()
	newVal := oldVal + i
	if newVal > define.REGISTER_CAPACITY {
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
	i := c.input.Value()
	c.a.setValue(i)
	c.carry.setValue(0)
}

func (c *CPU) mov_a(i define.Immidiate) {
	c.a.setValue(i)
	c.carry.setValue(0)
}

func (c *CPU) mov_ba() {
	im := c.a.value()
	c.b.setValue(im)
	c.carry.setValue(0)
}

func (c *CPU) add_b(i define.Immidiate) {
	oldVal := c.b.value()
	newVal := oldVal + i
	if newVal > define.REGISTER_CAPACITY {
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
	i := c.input.Value()
	c.b.setValue(i)
	c.carry.setValue(0)
}

func (c *CPU) mov_b(i define.Immidiate) {
	c.b.setValue(i)
	c.carry.setValue(0)
}

func (c *CPU) out_b() {
	i := c.b.value()
	c.output.SetValue(int(i))
	c.carry.setValue(0)
}

func (c *CPU) out(i define.Immidiate) {
	c.output.SetValue(int(i))
	c.carry.setValue(0)
}

func (c *CPU) jmp(i define.Immidiate) {
	c.pc.setValue(i)
	c.carry.setValue(0)
}

func (c *CPU) jnc(i define.Immidiate) {
	if c.carry.value() == 0 {
		c.pc.setValue(i)
	}
	c.carry.setValue(0)
}

func (c *CPU) execute(o Opecode, i define.Immidiate) error {
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

func (c *CPU) getPC() define.Adress {
	return define.Adress(c.pc.value())
}

func (c *CPU) progressPC() {
	v := c.pc.value()
	c.pc.setValue(v + 1)
}

func (c *CPU) Progress() {
	// fetch program instruction
	ad := c.getPC()
	inst := c.rom.Fetch(ad)

	// analyze fetch instruction
	ope, imm := c.decoder.Decode(inst)

	// execute opecode and immidiate
	err := c.execute(ope, imm)
	if err != nil {
		fmt.Println(err)
		return
	}

	c.progressPC()
}
