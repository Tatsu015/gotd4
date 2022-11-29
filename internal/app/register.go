package app

type Register struct {
	v int16
}

func NewRegister() Register {
	return Register{}
}

func (r *Register) value() int16 {
	return r.v
}

func (r *Register) setValue(v int16) {
	if v > REGISTER_CAPACITY {
		// when overflow set 0
		r.v = 0
	} else {
		r.v = v
	}
}
