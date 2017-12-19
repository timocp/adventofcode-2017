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
	tu.Walk()
	if tu.Letters != "ABCDEF" {
		t.Errorf("Walk() Letters => %s, want ABCDEF", tu.Letters)
	}
	if tu.Steps != 38 {
		t.Errorf("Walk() Letters => %d, want 38", tu.Steps)
	}
}
