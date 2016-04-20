package aca

import "testing"

func TestInsert(t *testing.T) {
	ac := NewACA()
	ac.Insert("say")
	ac.Insert("said")
	ac.Insert("shield")
	ac.Insert("shell")
	ac.PrintTree()
}
