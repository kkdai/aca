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
	// node *p = root;
	// int i=0,index;
	// while(str[i])
	// {
	//     index = str[i] - 'a';
	//     if(p->next[index] == NULL)
	//         p->next[index] = new node();
	//     p = p->next[index];
	//     i++;
	// }
	// p->count++;
	p := ac.root
	index := 0

	for i := 0; i < len(str); i++ {
		base := []byte("a")
		index = int(str[i] - base[0])
		if p.next[index] == nil {
			p.next[index] = NewNode()
		}
		p = p.next[index]
	}
	p.count++
}

//PrintTree :Print a tree
func (ac *ACA) PrintTree() {
	r := ac.root
	fmt.Printf("r->")
	recursiveTree(r)
}

func getString(index int) string {
	base := []byte("a")
	target := base[0] + byte(index)
	var str []byte
	str = append(str, target)
	return string(str)
}
func recursiveTree(current *Node) {
	for i := 0; i < CAPS; i++ {
		if current.next[i] != nil {
			fmt.Printf("%s->", getString(i))
			temp := current
			current = current.next[i]
			recursiveTree(current)
			current = temp
			for j := 0; j <= current.count; j++ {
				fmt.Printf("     ")
			}
		}
	}
	//tail node
	if current.count == 0 {
		fmt.Printf("null \n")
	}
}
