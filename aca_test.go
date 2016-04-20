package aca

import "testing"

func TestInsert(t *testing.T) {
	ac := NewACA()
	ac.Insert("say")
	ac.Insert("she")
	ac.PrintTree()
}
