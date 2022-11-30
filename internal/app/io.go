package app

type IO struct {
	v Immidiate
}

func NewIO() IO {
	return IO{}
}

func (r *IO) value() Immidiate {
	return r.v
}

func (r *IO) setValue(v Immidiate) {
	r.v = v & REGISTER_CAPACITY
}
