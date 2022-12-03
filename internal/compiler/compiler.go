package compiler

import (
	"bufio"
	"os"
)

type Compiler struct {
	codes []string
}

func NewCompiler(path string) Compiler {
	fp, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	codes := []string{}
	fScanner := bufio.NewScanner(fp)
	for fScanner.Scan() {
		codes = append(codes, fScanner.Text())
	}

	return Compiler{codes}
}

func (c *Compiler) Compile() []byte {
	b := []byte{0x01}
	ts := Analyze(c.codes)
	return b
}
