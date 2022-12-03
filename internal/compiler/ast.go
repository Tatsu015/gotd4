package compiler

func Parse(tokens []Token) Node {
	root := NewNode(0)
	for _, t := range tokens {
		root.add(NewNode(t.val))
	}
	return root
}
