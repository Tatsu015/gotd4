package app

type ROM struct {
	buf [16]byte
}

func NewROM() ROM {
	return ROM{}
}

func (r *ROM) Fetch(a Adress) Instruction {
	return Instruction(r.buf[a])
}
