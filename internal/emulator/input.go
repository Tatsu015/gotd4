package emulator

type Input struct {
	v Immidiate
}

func NewInput() Input {
	return Input{}
}

func (i *Input) value() Immidiate {
	return i.v
}

func (i *Input) setValue(v Immidiate) {
	i.v = v & REGISTER_CAPACITY
}
