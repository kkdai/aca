package aca

import "fmt"

//NewACA :Constructor
func NewACA() *ACA {
	ac := new(ACA)
	ac.root = NewNode()
	return ac
}

//ACA A-C Automation
type ACA struct {
	root *Node
	head int
	tail int
}

//Insert :Insert a new string into ACA
func (ac *ACA) Insert(str string) {
	p := ac.root

	for i := 0; i < len(str); i++ {
		index := getIndex(str[i])
		if p.next[index] == nil {
			p.next[index] = NewNode()
			p.count++
		}
		p = p.next[index]
	}
}

//PrintTree :Print a tree
func (ac *ACA) PrintTree() {
	r := ac.root
	fmt.Printf("r->")
	recursiveTree(r, 0)
}

func getIndex(char byte) int {
	base := []byte("a")
	return int(char - base[0])
}
func getString(index int) string {
	base := []byte("a")
	target := base[0] + byte(index)
	var str []byte
	str = append(str, target)
	return string(str)
}
func recursiveTree(current *Node, depth int) {
	for i := 0; i < CAPS; i++ {
		if current.next[i] != nil {
			fmt.Printf("%s->", getString(i))
			temp := current
			current = current.next[i]
			depth++
			recursiveTree(current, depth)
			current = temp
			fmt.Printf("(%d)\n", depth)
			for j := 0; j <= depth; j++ {
				fmt.Printf("    ")
			}
		}
	}
	//tail node
	if current.count == 0 {
		fmt.Printf("\tnull\n")
	}
}
