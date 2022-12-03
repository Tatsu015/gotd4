package io

import "github.com/Tatsu015/gotd4.git/internal/types"

type Input struct {
	v types.Immidiate
}

func NewInput() Input {
	return Input{}
}

func (i *Input) Value() types.Immidiate {
	return i.v
}

func (i *Input) SetValue(v types.Immidiate) {
	i.v = v & types.REGISTER_CAPACITY
}
