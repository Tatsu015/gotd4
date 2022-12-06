package cpu

import (
	"fmt"

	"github.com/Tatsu015/gotd4.git/internal/emulator/io"
	"github.com/Tatsu015/gotd4.git/internal/emulator/rom"
	"github.com/Tatsu015/gotd4.git/internal/types"
)

type CPU struct {
	a          Register
	b          Register
	carry      Register
	pc         Register
	decoder    Decoder
	rom        *rom.ROM
	input      *io.Input
	output     *io.Output
	currentOpe types.Opecode   // for debug display
	currentImm types.Immidiate // for debug display
}

func (c *CPU) Show() {
	ostr, _ := types.OpeToStr(c.currentOpe)
	fmt.Printf("a:%04b b:%04b pc:%04b carry:%04b in:%04b out:%04b ope:%v(%04b) imm:%04b\n",
		c.a.value(), c.b.value(), c.pc.value(), c.carry.value(), c.input.Value(), c.output.Value(), ostr, c.currentOpe, c.currentImm)
}

func NewCPU(rom *rom.ROM, input *io.Input, output *io.Output) CPU {
	return CPU{
		a:          NewRegister(),
		b:          NewRegister(),
		carry:      NewRegister(),
		pc:         NewRegister(),
		decoder:    NewDecoder(),
		rom:        rom,
		input:      input,
		output:     output,
		currentOpe: 0,
		currentImm: 0,
	}
}

func (c *CPU) add_a(i types.Immidiate) {
	oldVal := c.a.value()
	newVal := oldVal + i
	if newVal > types.REGISTER_CAPACITY {
		c.carry.setValue(1)
	} else {
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

func (c *CPU) mov_a(i types.Immidiate) {
	c.a.setValue(i)
	c.carry.setValue(0)
}

func (c *CPU) mov_ba() {
	im := c.a.value()
	c.b.setValue(im)
	c.carry.setValue(0)
}

func (c *CPU) add_b(i types.Immidiate) {
	oldVal := c.b.value()
	newVal := oldVal + i
	if newVal > types.REGISTER_CAPACITY {
		c.carry.setValue(1)
	} else {
		c.carry.setValue(0)
	}
	c.b.setValue(newVal)
}

func (c *CPU) in_b() {
	i := c.input.Value()
	c.b.setValue(i)
	c.carry.setValue(0)
}

func (c *CPU) mov_b(i types.Immidiate) {
	c.b.setValue(i)
	c.carry.setValue(0)
}

func (c *CPU) out_b() {
	i := c.b.value()
	c.output.SetValue(int(i))
	c.carry.setValue(0)
}

func (c *CPU) out(i types.Immidiate) {
	c.output.SetValue(int(i))
	c.carry.setValue(0)
}

func (c *CPU) jmp(i types.Immidiate) {
	c.pc.setValue(i)
	c.carry.setValue(0)
}

func (c *CPU) jnc(i types.Immidiate) {
	if c.carry.value() == 0 {
		c.pc.setValue(i)
	}
	c.carry.setValue(0)
}

func (c *CPU) execute(o types.Opecode, i types.Immidiate) error {
	switch o {
	case types.ADD_A:
		c.add_a(i)
		return nil
	case types.MOV_AB:
		c.mov_ab()
		return nil
	case types.IN_A:
		c.in_a()
		return nil
	case types.MOV_A:
		c.mov_a(i)
		return nil
	case types.MOV_BA:
		c.mov_ba()
		return nil
	case types.ADD_B:
		c.add_b(i)
		return nil
	case types.IN_B:
		c.in_b()
		return nil
	case types.MOV_B:
		c.mov_b(i)
		return nil
	case types.OUT_B:
		c.out_b()
		return nil
	case types.OUT:
		c.out(i)
		return nil
	case types.JMP:
		c.jmp(i)
		return nil
	case types.JNC:
		c.jnc(i)
		return nil
	default:
		return fmt.Errorf("opecode %v not exist!", o)
	}
}

func (c *CPU) getPC() types.Adress {
	return types.Adress(c.pc.value())
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
	c.currentOpe, c.currentImm = c.decoder.Decode(inst)

	// execute opecode and immidiate
	err := c.execute(c.currentOpe, c.currentImm)
	if err != nil {
		fmt.Println(err)
		return
	}

	c.progressPC()
}
