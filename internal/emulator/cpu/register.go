package cpu

import "github.com/Tatsu015/gotd4.git/internal/define"

type Register struct {
	v define.Immidiate
}

func NewRegister() Register {
	return Register{}
}

func (r *Register) value() define.Immidiate {
	return r.v
}

func (r *Register) setValue(v define.Immidiate) {
	r.v = v & define.REGISTER_CAPACITY
}
