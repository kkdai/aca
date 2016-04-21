//The MIT License (MIT)
//Copyright (c) 2015 Evan Lin
//Permission is hereby granted, free of charge, to any person obtaining a copy
//of this software and associated documentation files (the "Software"), to deal
//in the Software without restriction, including without limitation the rights
//to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//copies of the Software, and to permit persons to whom the Software is
//furnished to do so, subject to the following conditions:
//The above copyright notice and this permission notice shall be included in all
//copies or substantial portions of the Software.
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
//SOFTWARE.

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

	nodeList  [5000]*Node // For failed function constructor
	lineBreak bool
}

//Insert :Insert a new string into ACA
func (ac *ACA) Insert(str string) {
	p := ac.root

	for i := 0; i < len(str); i++ {
		index := getIndex(str[i])
		if p.next[index] == nil {
			p.next[index] = NewNode()
		}
		p = p.next[index]
	}
	p.end = true
}

//PrintTree :Print a tree
func (ac *ACA) PrintTree() {
	r := ac.root
	fmt.Printf("r->")
	ac.recursiveTree(r, 0)
}

func (ac *ACA) buildAC() {
	r := ac.root
	r.fail = nil
	ac.nodeList[ac.head] = r
	ac.head++

	for {
		if ac.head == ac.tail {
			break
		}

		temp := ac.nodeList[ac.tail]
		ac.tail++
		var p *Node

		for i := 0; i < CAPS; i++ {
			if temp.next[i] != nil {
				if temp == ac.root {
					temp.next[i].fail = ac.root
				} else {
					p = temp.fail
					for {
						if p == nil {
							break
						}

						if p.next[i] != nil {
							temp.next[i].fail = p.next[i]
							break
						}
						p = p.fail
					}

					if p == nil {
						temp.next[i].fail = r
					}
				}

				ac.nodeList[ac.head] = temp.next[i]
				ac.head++
			}
		}

	}
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

func (ac *ACA) recursiveTree(current *Node, depth int) {
	for i := 0; i < CAPS; i++ {
		if current.next[i] != nil {
			if ac.lineBreak == true {
				for j := 0; j < depth; j++ {
					fmt.Printf("   ")
				}
				ac.lineBreak = false
			}

			fmt.Printf("%s->", getString(i))
			temp := current
			current = current.next[i]
			depth++
			ac.recursiveTree(current, depth)
			current = temp
			ac.lineBreak = true
		}
	}

	//tail node
	if current.end == true {
		fmt.Printf("null\n")
	}
}
