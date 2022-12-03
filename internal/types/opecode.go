package types

import "fmt"

type Opecode int16

const (
	ADD_A  Opecode = 0x00
	MOV_AB Opecode = 0x01
	IN_A   Opecode = 0x02
	MOV_A  Opecode = 0x03
	MOV_BA Opecode = 0x04
	ADD_B  Opecode = 0x05
	IN_B   Opecode = 0x06
	MOV_B  Opecode = 0x07
	OUT_B  Opecode = 0x09
	OUT    Opecode = 0x0b
	JMP    Opecode = 0x0e
	JNC    Opecode = 0x0f
)

func codeMap() map[string]Opecode {
	CODE_MAP := map[string]Opecode{
		"ADD A":   ADD_A,
		"MOV A,B": MOV_AB,
		"IN A":    IN_A,
		"MOV A":   MOV_A,
		"MOV B,A": MOV_BA,
		"ADD B":   ADD_B,
		"IN B":    IN_B,
		"MOV B":   MOV_B,
		"OUT B":   OUT_B,
		"OUT":     OUT,
		"JMP":     JMP,
		"JNC":     JNC,
	}
	return CODE_MAP
}

func IsOpe(s string) bool {
	_, ok := codeMap()[s]
	return ok
}

func StrToOpe(s string) (Opecode, error) {
	if ope, ok := codeMap()[s]; ok {
		return 0, fmt.Errorf("Error: Not known opecode. %v", s)
	} else {
		return ope, nil
	}
}
