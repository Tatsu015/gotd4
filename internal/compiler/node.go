package compiler

import "fmt"

type Node struct {
	val      int
	children []Node
}

func NewNode(val int) Node {
	return Node{val, []Node{}}
}

func (n *Node) add(c Node) {
	n.children = append(n.children, c)
}

func (n *Node) ntob() []byte {
	t := []byte{}
	for _, c := range n.children {
		b := c.ntob()
		t = append(t, b...)
	}
	return append(t, byte(n.val))
}

func (n *Node) squash(b []byte) []byte {
	// TD4 is 4bit CPU
	// but, golang has no bit datatype. minimum unit is 1byte(8bit).
	// Now opecode and immidiate size both 8 bit in this program,
	// so simply parse AST, one instruction(opecode + immidiate) become 2byte(16bit).
	// that is error!
	// so we must to squash instruction from 16bit to 8bit
	size := len(b)
	if size%2 != 0 {
		e := fmt.Sprintf("Error: immidiate size must be 8byte.")
		panic(e)
	}

	squashed := []byte{}
	for i := 0; i < size/2; i++ {
		upper := b[2*i]
		lower := b[2*i+1]
		ins := upper<<4 | lower
		squashed = append(squashed, ins)
	}
	return squashed
}

func (n *Node) toBinary() []byte {
	b := n.ntob()
	b = b[:len(b)-1] // remove root element
	s := n.squash(b)
	return s
}
