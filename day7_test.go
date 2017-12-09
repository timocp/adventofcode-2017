package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

var sampleProgram = []string{
	"pbga (66)",
	"xhth (57)",
	"ebii (61)",
	"havc (66)",
	"ktlj (57)",
	"fwft (72) -> ktlj, cntj, xhth",
	"qoyq (66)",
	"padx (45) -> pbga, havc, qoyq",
	"tknk (41) -> ugml, padx, fwft",
	"jptl (61)",
	"ugml (68) -> gyxo, ebii, jptl",
	"gyxo (61)",
	"cntj (57)",
}

func TestParseNode(t *testing.T) {
	for _, tt := range []struct {
		in        string
		name      string
		weight    int
		balancing []string
	}{
		{sampleProgram[0], "pbga", 66, []string{}},
		{sampleProgram[1], "xhth", 57, []string{}},
		{sampleProgram[2], "ebii", 61, []string{}},
		{sampleProgram[3], "havc", 66, []string{}},
		{sampleProgram[4], "ktlj", 57, []string{}},
		{sampleProgram[5], "fwft", 72, []string{"ktlj", "cntj", "xhth"}},
		{sampleProgram[6], "qoyq", 66, []string{}},
		{sampleProgram[7], "padx", 45, []string{"pbga", "havc", "qoyq"}},
		{sampleProgram[8], "tknk", 41, []string{"ugml", "padx", "fwft"}},
		{sampleProgram[9], "jptl", 61, []string{}},
		{sampleProgram[10], "ugml", 68, []string{"gyxo", "ebii", "jptl"}},
		{sampleProgram[11], "gyxo", 61, []string{}},
		{sampleProgram[12], "cntj", 57, []string{}},
	} {
		p, err := ParseNode(tt.in)
		if err != nil {
			t.Errorf("ParseNode(%s) returned error %v, expected nil", tt.in, err)
		} else {
			if p.Name != tt.name {
				t.Errorf("ParseNode(%s).Name => %s, want %s", tt.in, p.Name, tt.name)
			}
			if p.Weight != tt.weight {
				t.Errorf("ParseNode(%s).Weight => %d, want %d", tt.in, p.Weight, tt.weight)
			}
			if !slicesEqual(p.Balancing, tt.balancing) {
				t.Errorf("ParseNode(%s).Balancing => %q, want %q", tt.in, p.Balancing, tt.balancing)
			}
		}
	}
}

func TestReadTower(t *testing.T) {
	r, err := ReadTower(bytes.NewBufferString(strings.Join(sampleProgram, "\n")))
	if err != nil {
		t.Errorf("ReadTower() returned error %v, want nil", err)
	} else if len(r) != len(sampleProgram) {
		t.Errorf("ReadTower() length => %d, want %d", len(r), len(sampleProgram))
	}
}

func TestFindBottomNode(t *testing.T) {
	tower, _ := ReadTower(bytes.NewBufferString(strings.Join(sampleProgram, "\n")))
	r := tower.FindBottomNode()
	exp := "tknk"
	if r == nil {
		t.Errorf("FindBottomNode returned nil")
	} else if r.Name != exp {
		t.Errorf("FindBottomNode() => %s, want %s", r.Name, exp)
	}
}

func slicesEqual(a []string, b []string) bool {
	if len(a) != len(b) {
		fmt.Printf("%d, %d\n", len(a), len(b))
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
