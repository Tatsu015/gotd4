package io

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/Tatsu015/gotd4.git/internal/types"
)

type Output struct {
	v int
}

func NewOutput() Output {
	return Output{}
}

func (o *Output) replace01(org string, zero string, one string) string {
	t1 := strings.ReplaceAll(org, "0", zero)
	t2 := strings.ReplaceAll(t1, "1", one)
	return t2
}

func (o *Output) InitDisplay() {
	o.reset()
	o.firstShow()
}

func (o *Output) reset() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func (o *Output) firstShow() {
	fmt.Println("□□□□")
}

func (o *Output) Show() {
	o.reset()
	bitStr := fmt.Sprintf("%04b", o.v)
	disp := o.replace01(bitStr, "□", "■")
	fmt.Println(disp)
}

func (o *Output) Value() int {
	return o.v
}

func (o *Output) SetValue(v int) {
	o.v = v & types.REGISTER_CAPACITY
}
