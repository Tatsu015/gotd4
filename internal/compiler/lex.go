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

func divide(line string) (opecode string, immidiate string, err error) {
	l := strings.TrimSpace(line)
	if strings.Contains(l, ",") {
		ss := strings.Split(l, ",")
		opecode = simplfy(ss[0])
		immidiate = simplfy(ss[1])
		return opecode, immidiate, nil
	}
	if strings.Contains(l, " ") {
		ss := strings.Split(l, " ")
		opecode = simplfy(ss[0])
		immidiate = simplfy(ss[1])
		return opecode, immidiate, nil
	}
	return "", "", fmt.Errorf("Error: failed to divide opecode and immidiate.")
}

func Analyze(codes []string) []Token {
	tokens := []Token{}
	for i, line := range codes {
		// line has 2 pattern, contain ',' or not.
		// for ex.
		//  ADD A, 0011 <- contain ','
		//  IN A <- not contain ','
		// if contain ',', Opecode is before ',' and Immidiate is after ','.
		// if not contain ',', Opecode is before ' ' and Immidiate is after ' '.
		ope, imm, err := divide(line)
		if err != nil {
			err := fmt.Sprintf("Error: syntax error at l:%d %s", i, line)
			panic(err)
		}

		i, err := types.StrtoImm(imm)
		fmt.Println(ope, "-", i)

	}
	return tokens
}
