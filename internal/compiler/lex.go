package compiler

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/Tatsu015/gotd4.git/internal/types"
)

func simplfy(s string) string {
	re := regexp.MustCompile("[\\s\n\r]")
	t := re.ReplaceAll([]byte(s), []byte{})
	return string(t)
}

func divide(line string) (opecode string, immidiate string, hasImm bool, err error) {
	l := strings.TrimSpace(line)
	if types.IsOpe(l) {
		return l, "", false, nil
	}
	if strings.Contains(l, ",") {
		ss := strings.Split(l, ",")
		opecode = simplfy(ss[0])
		immidiate = simplfy(ss[1])
		return opecode, immidiate, true, nil
	}
	if strings.Contains(l, " ") {
		ss := strings.Split(l, " ")
		opecode = simplfy(ss[0])
		immidiate = simplfy(ss[1])
		return opecode, immidiate, true, nil
	}
	return "", "", false, fmt.Errorf("Error: failed to divide opecode and immidiate.")
}

func Analyze(codes []string) []Token {
	tokens := []Token{}
	for i, line := range codes {
		// line has 3 pattern, contain ',' or not.
		// for ex.
		// 1. MOV A,B <- no immidiate
		// 2. ADD A, 0011 <- contain ','
		// 3. IN A <- not contain ','
		// if contain ',', Opecode is before ',' and Immidiate is after ','.
		// if not contain ',', Opecode is before ' ' and Immidiate is after ' '.
		ope, imm, hasImm, err := divide(line)
		if err != nil {
			e := fmt.Sprintf("Error: syntax error at l:%d %s", i, line)
			panic(e)
		}

		o, err := types.StrToOpe(ope)
		if err != nil {
			e := fmt.Sprintf("Error: syntax error at l:%d %s", i, line)
			panic(e)
		}
		tokens = append(tokens, NewToken(Ope, int(o)))

		// exist opecode only
		if hasImm == false {
			continue
		}

		i, err := types.StrtoImm(imm)
		if err != nil {
			e := fmt.Sprintf("Error: syntax error at l:%d %s", i, line)
			panic(e)
		}
		tokens = append(tokens, NewToken(Imm, int(i)))
	}
	return tokens
}
