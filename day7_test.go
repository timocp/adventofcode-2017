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
	} else {
		if r.Name != "tknk" {
			t.Errorf("r.Name => %s, want %s", r.Name, "tknk")
		}
	}
}

func TestIterate(t *testing.T) {
	tower, _ := ReadTower(bytes.NewBufferString(strings.Join(sampleProgram, "\n")))
	names := []string{}
	tower.Iterate(func(n *Node, depth int) {
		names = append(names, fmt.Sprintf("%d:%s", depth, n.Name))
	})
	exp := []string{
		"0:tknk",
		"1:ugml", "2:gyxo", "2:ebii", "2:jptl",
		"1:padx", "2:pbga", "2:havc", "2:qoyq",
		"1:fwft", "2:ktlj", "2:cntj", "2:xhth",
	}
	if !slicesEqual(names, exp) {
		t.Errorf("Iterate Names => %q, want %q", names, exp)
	}
}

func TestWeight(t *testing.T) {
	tower, _ := ReadTower(bytes.NewBufferString(strings.Join(sampleProgram, "\n")))
	r := tower.TotalWeight()
	if r != 778 {
		t.Errorf("TotalWeight() => %d, want %d", r, 778)
	}
}

func TestWrongWeightShouldBe(t *testing.T) {
	tower, _ := ReadTower(bytes.NewBufferString(strings.Join(sampleProgram, "\n")))
	r := tower.WrongWeightShouldBe()
	exp := 60
	if r != exp {
		t.Errorf("WrongWeightShouldBe() => %d, want %d", r, exp)
	}
}

func slicesEqual(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
