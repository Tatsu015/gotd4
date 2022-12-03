package cpu

import "github.com/Tatsu015/gotd4.git/internal/types"

type Register struct {
	v types.Immidiate
}

func NewRegister() Register {
	return Register{}
}

func (r *Register) value() types.Immidiate {
	return r.v
}

func (r *Register) setValue(v types.Immidiate) {
	r.v = v & types.REGISTER_CAPACITY
}
