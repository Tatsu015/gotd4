package compiler

type Kind int

const (
	Ope = 0
	Imm = 1
)

type Token struct {
	kind Kind
	val  int
}

func NewToken(kind Kind, val int) Token {
	return Token{kind, val}
}
