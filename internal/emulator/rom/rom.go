package rom

import "github.com/Tatsu015/gotd4.git/internal/types"

type ROM struct {
	program []byte
}

func NewROM(program []byte) ROM {
	return ROM{
		program: program,
	}
}

func (r *ROM) MaxCapacityByte() int {
	return types.REGISTER_CAPACITY
}

func (r *ROM) Fetch(a types.Adress) types.Instruction {
	return types.Instruction(r.program[a])
}
