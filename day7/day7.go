package day7

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

// Node ...
type Node struct {
	Name      string
	Weight    int
	Balancing []string
	Parent    *Node
	Children  []*Node
}

var nodeRe = regexp.MustCompile(`^(\w+) \((\d+)\)(?: -> )?(.*)$`)

// ParseNode reads a node definition from a single line.  Balancing contains
// a list of names of nodes it is balancing, these will be used to populate
// Children[] and Parent
func ParseNode(input string) (*Node, error) {
	if !nodeRe.MatchString(input) {
		return nil, fmt.Errorf("ParseNode: Invalid format: %s", input)
	}
	matches := nodeRe.FindStringSubmatch(input)
	weight, err := strconv.Atoi(matches[2])
	if err != nil {
		return nil, err
	}
	return &Node{
		Name:      matches[1],
		Weight:    weight,
		Balancing: splitWords(matches[3]),
	}, nil
}

// ReadTower builds a set of nodes from input.  The top node is returned.
func ReadTower(input io.Reader) (*Node, error) {
	// a hash of nodes keyed by their names
	tower := make(map[string]*Node)
	s := bufio.NewScanner(input)
	for s.Scan() {
		n, err := ParseNode(s.Text())
		if err != nil {
			return nil, err
		}
		tower[n.Name] = n
	}
	if s.Err() != nil {
		return nil, s.Err()
	}
	// build parent/child relationships
	for _, node := range tower {
		for _, childName := range node.Balancing {
			node.Children = append(node.Children, tower[childName])
			tower[childName].Parent = node
		}
	}
	// find the node with no parent, it is the root
	for _, node := range tower {
		if node.Parent == nil {
			return node, nil
		}
	}
	return nil, fmt.Errorf("Tree has no root node")
}

// Iterate over a tree, calling f(node, depth) for each node
func (node *Node) Iterate(f func(*Node, int)) {
	_iterate(node, 0, f)
}

func _iterate(node *Node, depth int, f func(*Node, int)) {
	f(node, depth)
	for _, child := range node.Children {
		_iterate(child, depth+1, f)
	}
}

// TotalWeight returns the weight, including of all child nodes
func (node *Node) TotalWeight() int {
	weight := 0
	node.Iterate(func(n *Node, depth int) {
		weight += n.Weight
	})
	return weight
}

func (node *Node) String() string {
	s := ""
	node.Iterate(func(n *Node, depth int) {
		for i := 0; i < depth; i++ {
			s += " "
		}
		s += fmt.Sprintf("%s (%d)\n", n.Name, n.TotalWeight())
	})
	return s
}

// WrongWeightShouldBe returns the wight that the incorrectly balanced node
// should have (given that exactly one program is the wrong weight)
//
// this means finding a node where its childrens weights are not all the same,
// and returning what the wrong child's weight should be.
func (node *Node) WrongWeightShouldBe() int {
	correctWeight := 0
	wrongDepth := -1
	node.Iterate(func(n *Node, depth int) {
		if len(n.Children) == 0 {
			return
		}
		weights := make([]int, len(n.Children))
		for i, child := range n.Children {
			weights[i] = child.TotalWeight()
		}
		// work out the more common number
		common := findCommon(weights)
		// any child weights which don't match the expected value?
		for i, child := range n.Children {
			if weights[i] != common {
				// this is the node that is wrong; we want the one at the
				// maximum depth
				if depth > wrongDepth {
					correctWeight = child.Weight - (weights[i] - common)
					wrongDepth = depth
				}
			}
		}
	})
	return correctWeight
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

// in a list with at most 2 different unique numbers, return the one which
// occurs more than once
func findCommon(numbers []int) int {
	a := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if numbers[i] != a {
			if i == len(numbers)-1 {
				return a
			} else if numbers[i+1] == a {
				return a
			} else {
				return numbers[i]
			}
		}
	}
	return a
}
