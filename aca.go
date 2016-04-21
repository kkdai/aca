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

import (
	"fmt"
	"strings"
)

//NewACA :Constructor
func NewACA() *ACA {
	ac := new(ACA)
	ac.root = NewNode(ac.nodeCount)
	ac.nodeCount++
	return ac
}

//ACA A-C Automation
type ACA struct {
	root        *Node
	head        int
	tail        int
	nodeCount   int
	stringCount int
	stringList  []string

	nodeList  [5000]*Node // For failed function constructor
	lineBreak bool
}

//Insert :Insert a new string into ACA
func (ac *ACA) Insert(oriStr string) {
	str := strings.ToLower(oriStr)
	p := ac.root
	ac.stringCount++

	for i := 0; i < len(str); i++ {
		index := getIndex(str[i])
		if p.next[index] == nil {
			p.next[index] = NewNode(ac.nodeCount)
			ac.nodeCount++
		}
		p = p.next[index]
	}
	p.strNo = ac.stringCount

	ac.stringList = append(ac.stringList, oriStr)
}

//PrintTree :Print a tree
func (ac *ACA) PrintTree() {
	r := ac.root
	fmt.Printf("r(0)->")
	ac.recursivePrintNode(r, 0)
}

//BuildAC :Build this AC Automation failed function according to current tree
func (ac *ACA) BuildAC() {
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
					// temp.fail = ac.root
				} else {
					p = temp.fail
					for {
						if p == nil {
							break
						}

						// if current's next node fail is equal to your fail's next
						// assign current's next failed to your failed's next.
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

//Query : Return cnt (hit count) for haystack include how many needles
func (ac *ACA) Query(haystack string) []string {
	haystack = strings.ToLower(haystack)

	n := len(haystack)
	var index int
	var retSlice []string

	p := ac.root

	for i := 0; i < n; i++ {
		index = getIndex(haystack[i])
		for {
			if p.next[index] == nil && p != ac.root {
				p = p.fail
			} else {
				break
			}
		}

		p = p.next[index]
		if p == nil {
			p = ac.root
		}

		temp := p

		for {
			if temp == ac.root || temp.strNo == 0 {
				break
			}
			// cnt = cnt + temp.strNo
			if temp.strNo > 0 {
				//strNo is 1-base, but index is zero-based
				stringIndex := temp.strNo - 1
				fmt.Println("Found string:", ac.stringList[stringIndex], " in i=", i)
				retSlice = append(retSlice, ac.stringList[stringIndex])
			}
			temp = temp.fail
		}
	}
	return retSlice
}

func (ac *ACA) recursivePrintNode(current *Node, depth int) {
	for i := 0; i < CAPS; i++ {
		if current.next[i] != nil {
			if ac.lineBreak == true {
				for j := 0; j < depth; j++ {
					fmt.Printf("          ")
				}
				ac.lineBreak = false
			}

			var failNodeID int
			if current.next[i].fail != nil && current.fail != nil {
				failNodeID = current.fail.id
			}
			fmt.Printf("%s( %02d,%02d [%d] )->", getString(i), current.next[i].id, failNodeID, current.next[i].strNo)
			temp := current
			current = current.next[i]
			depth++
			ac.recursivePrintNode(current, depth)
			current = temp
			ac.lineBreak = true
		}
	}

	//tail node
	if current.strNo > 0 {
		fmt.Printf(" null\n")
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
