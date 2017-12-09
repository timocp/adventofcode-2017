package main

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"
)

// Node ...
type Node struct {
	Name      string
	Weight    int
	Balancing []string
	HeldBy    string
}

// Tower is a set of nodes, keyed by their name
type Tower map[string]*Node

var nodeRe = regexp.MustCompile(`^(\w+) \((\d+)\)(?: -> )?(.*)$`)

// ParseNode reads a node definition from a single line
func ParseNode(input string) (*Node, error) {
	if !nodeRe.MatchString(input) {
		return nil, fmt.Errorf("ParseNode: Invalid format: %s", input)
	}
	matches := nodeRe.FindStringSubmatch(input)
	return &Node{
		Name:      matches[1],
		Weight:    toInt(matches[2]),
		Balancing: splitWords(matches[3]),
	}, nil
}

// ReadTower builds a Tower of nodes from input
func ReadTower(input io.Reader) (Tower, error) {
	tower := make(map[string]*Node)
	s := bufio.NewScanner(input)
	for s.Scan() {
		n, err := ParseNode(s.Text())
		if err != nil {
			return tower, nil
		}
		tower[n.Name] = n
	}
	// record parents of each node ("" = none)
	for name, node := range tower {
		for _, child := range node.Balancing {
			tower[child].HeldBy = name
		}
	}
	return tower, s.Err()
}

// FindBottomNode returns the node at the bottom of the tower
// The bottom node is the one which isn't being held up by any other
func (tower Tower) FindBottomNode() *Node {
	for _, node := range tower {
		if node.HeldBy == "" {
			return node
		}
	}
	return nil
}

// split a string by commas and remove whitespace
func splitWords(s string) []string {
	if strings.TrimSpace(s) == "" {
		return []string{}
	}
	words := strings.Split(s, ",")
	for i := range words {
		words[i] = strings.TrimSpace(words[i])
	}
	return words
}
