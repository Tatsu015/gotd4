package rom

import "github.com/Tatsu015/gotd4.git/internal/define"

type ROM struct {
	program []byte
}

func NewROM(program []byte) ROM {
	return ROM{
		program: program,
	}
}

func (r *ROM) MaxCapacityByte() int {
	return define.REGISTER_CAPACITY
}

func (r *ROM) Fetch(a define.Adress) define.Instruction {
	return define.Instruction(r.program[a])
}
