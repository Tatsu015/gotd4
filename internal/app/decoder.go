package app

type Decoder struct{}

func NewDecoder() Decoder {
	return Decoder{}
}

func (d *Decoder) Decode(i Instruction) (Opecode, Immidiate) {
	return 1, 1 // TODO
}
