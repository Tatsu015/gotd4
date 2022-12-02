package emulator

import "fmt"

type Output struct {
	v Immidiate
}

func NewOutput() Output {
	return Output{}
}

func (o *Output) Show() {
	fmt.Printf("%04b\n", o.v)
}

func (o *Output) value() Immidiate {
	return o.v
}

func (o *Output) setValue(v Immidiate) {
	o.v = v & REGISTER_CAPACITY
}
