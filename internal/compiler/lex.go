package compiler

import (
	"fmt"
	"regexp"
	"strings"
)

func simplfy(s string) string {
	re := regexp.MustCompile("[\\s\n\r]")
	t := re.ReplaceAll([]byte(s), []byte{})
	return string(t)
}

func Analyze(codes []string) []Token {
	tokens := []Token{}
	for i, code := range codes {
		ss := strings.Split(code, ",")
		// opecode divided by ','
		if len(ss) == 2 {
			ope := simplfy(ss[0])
			tokens = append(tokens, t0)
			imm := simplfy(ss[1])
			continue
		}
		ss = strings.Split(code, " ")
		if len(ss) == 2 {
			ope := simplfy(ss[0])
			tokens = append(tokens, t0)
			imm := simplfy(ss[1])
			continue
		}
		err := fmt.Sprintf("Error: syntax error at l:%d %s", i, code)
		panic(err)
	}
	return []Token{}
}
