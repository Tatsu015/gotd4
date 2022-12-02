package emulator

type Register struct {
	v Immidiate
}

func NewRegister() Register {
	return Register{}
}

func (r *Register) value() Immidiate {
	return r.v
}

func (r *Register) setValue(v Immidiate) {
	r.v = v & REGISTER_CAPACITY
}
