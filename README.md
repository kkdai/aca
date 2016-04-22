ACA: Aho–Corasick Automation
==================

[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/kkdai/aca/master/LICENSE)  [![GoDoc](https://godoc.org/github.com/kkdai/aca?status.svg)](https://godoc.org/github.com/kkdai/aca)  [![Build Status](https://travis-ci.org/kkdai/aca.svg?branch=master)](https://travis-ci.org/kkdai/aca)


### Aho–Corasick algorithm

In computer science, the Aho–Corasick algorithm is a string searching algorithm invented by Alfred V. Aho and Margaret J. Corasick.[1] It is a kind of dictionary-matching algorithm that locates elements of a finite set of strings (the "dictionary") within an input text. It matches all strings simultaneously. The complexity of the algorithm is linear in the length of the strings plus the length of the searched text plus the number of output matches. Note that because all matches are found, there can be a quadratic number of matches if every substring matches (e.g. dictionary = a, aa, aaa, aaaa and input string is aaaa).

(from [Wiki](https://en.wikipedia.org/wiki/Aho%E2%80%93Corasick_algorithm))

## Explanation in Chinese

Check [this blog](http://www.evanlin.com/about-aca/) for more detail explanation, but it is Chinese
 
Install
---------------
`go get github.com/kkdai/aca`


Usage
---------------

```go
package main

import (
	"fmt"

	. "github.com/kkdai/aca"
)

func main() {
	ac := NewACA()
	ac.Insert("say")
	ac.Insert("she")
	ac.Insert("shell")
	ac.Insert("shr")
	ac.Insert("her")
	ac.BuildAC()
	fmt.Println(ac.Query("aaashellaashrmmmmmhemmhera"))
	//[shell, shr, her]
}
```


Inspired
---------------

- [跳跃表，字典树（单词查找树，Trie树），后缀树，KMP算法，AC 自动机相关算法原理详细汇总](http://blog.csdn.net/zhongwen7710/article/details/39274349)
- [Biosequence Algorithms, Spring 2005 Lecture 4: Set Matching and Aho-Corasick Algorithm](http://www.cs.uku.fi/~kilpelai/BSA05/lectures/slides04.pdf)
- [Wiki: Aho–Corasick algorithm](https://en.wikipedia.org/wiki/Aho%E2%80%93Corasick_algorithm)

Project52
---------------

It is one of my [project 52](https://github.com/kkdai/project52).


License
---------------

This package is licensed under MIT license. See LICENSE for details.

