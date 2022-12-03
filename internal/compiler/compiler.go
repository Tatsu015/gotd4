package compiler

import (
	"bufio"
	"fmt"
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
	fmt.Println("---Compile---")
	b := []byte{0x01}
	return b
}
