package cpu

import "github.com/Tatsu015/gotd4.git/internal/define"

type Decoder struct{}

func NewDecoder() Decoder {
	return Decoder{}
}

func (d *Decoder) Decode(i define.Instruction) (Opecode, define.Immidiate) {
	imm := i & 0x0f
	ope := (i & 0xf0) >> 4

	return Opecode(ope), define.Immidiate(imm)
}
