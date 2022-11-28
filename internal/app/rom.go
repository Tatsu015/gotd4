package app

type ROM struct{}

func NewROM() ROM {
	return ROM{}
}

func (r *ROM) Fetch() Instruction {
	return [4]byte{1} // TODO
}
