package main

// Issue #2 (https://github.com/kyoh86/exportloopref/issues/2
// This is valid code because n takes new pointer values, so &n.name points correctly to different names.

import (
	"fmt"
)

type Node struct {
	name string
}

type Nodes []*Node

type Graph struct {
	nodes Nodes
}

func main() {
	var graph Graph
	var s *string

	graph.nodes = Nodes{&Node{"hello"}, &Node{"world"}, &Node{"foo"}, &Node{"bar"}, &Node{"baz"}}

	for i, n := range graph.nodes {
		if i == 2 {
			s = &n.name // here
		}
	}

	fmt.Println(*s)
}
