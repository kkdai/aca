package aca

const (
	CAPS int = 26
)

type Node struct {
	count int
	fail  *Node
	next  [CAPS]*Node
}

func NewNode() *Node {
	n := new(Node)
	return n
}
