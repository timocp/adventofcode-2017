package day19

import "testing"

var in = `     |          
     |  +--+    
     A  |  C    
 F---|----E|--+ 
     |  |  |  D 
     +B-+  +--+ 
`

func TestWalk(t *testing.T) {
	tu := NewTubes(in)
	if tu.col != 5 {
		t.Errorf("NewTubes().col => %d, want %d", tu.col, 5)
	}
	r := tu.Walk()
	exp := "ABCDEF"
	if r != exp {
		t.Errorf("WalkLetters() => %s, want %s", r, exp)
	}
}
