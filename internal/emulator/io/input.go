package io

import "github.com/Tatsu015/gotd4.git/internal/define"

type Input struct {
	v define.Immidiate
}

func NewInput() Input {
	return Input{}
}

func (i *Input) Value() define.Immidiate {
	return i.v
}

func (i *Input) SetValue(v define.Immidiate) {
	i.v = v & define.REGISTER_CAPACITY
}
