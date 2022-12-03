package cpu

import "github.com/Tatsu015/gotd4.git/internal/types"

type Decoder struct{}

func NewDecoder() Decoder {
	return Decoder{}
}

func (d *Decoder) Decode(i types.Instruction) (Opecode, types.Immidiate) {
	imm := i & 0x0f
	ope := (i & 0xf0) >> 4

	return Opecode(ope), types.Immidiate(imm)
}
