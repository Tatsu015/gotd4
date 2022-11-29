package app

type ROM struct {
	program []byte
}

func NewROM(program []byte) ROM {
	return ROM{
		program: program,
	}
}

func (r *ROM) MaxCapacityByte() int {
	return ROM_CAPACITY
}

func (r *ROM) Fetch(a Adress) Instruction {
	return Instruction(r.program[a])
}
