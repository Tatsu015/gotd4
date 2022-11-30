package app

type Decoder struct{}

func NewDecoder() Decoder {
	return Decoder{}
}

func (d *Decoder) Decode(i Instruction) (Opecode, Immidiate) {
	imm := i & 0x0f
	ope := (i & 0xf0) >> 4

	return Opecode(ope), Immidiate(imm)
}
