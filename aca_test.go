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

package aca_test

import (
	"fmt"
	"testing"

	. "github.com/kkdai/aca"
)

func TestEmpty(t *testing.T) {
	ac := NewACA()
	ac.BuildAC()
}

func TestInsert(t *testing.T) {
	ac := NewACA()
	ac.Insert("say")
	ac.Insert("she")
	ac.Insert("shell")
	ac.Insert("shr")
	ac.Insert("her")
	ac.BuildAC()
	ret := ac.Query("aaashellaashrmmmmmhemmhera")
	if len(ret) != 4 {
		t.Error("Query error")
	}
	fmt.Println("find string:", ret)
}
