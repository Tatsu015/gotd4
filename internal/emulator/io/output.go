package io

import (
	"fmt"

	"github.com/Tatsu015/gotd4.git/internal/define"
)

type Output struct {
	v int
}

func NewOutput() Output {
	return Output{}
}

func (o *Output) Show() {
	fmt.Printf("%04b\n", o.v)
}

func (o *Output) Value() int {
	return o.v
}

func (o *Output) SetValue(v int) {
	o.v = v & define.REGISTER_CAPACITY
}
