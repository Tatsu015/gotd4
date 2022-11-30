package app

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
	if v > REGISTER_CAPACITY {
		// when overflow set 0
		r.v = 0
	} else {
		r.v = v
	}
}
