package main

import (
	"fmt"

	"gonum.org/v1/gonum/graph/encoding/dot"
	"gonum.org/v1/gonum/graph/simple"
)

type Node struct {
	id int64
}

func (n Node) ID() int64 {
	return n.id
}

func main() {
	g := simple.NewDirectedGraph()

	// Add vertices to the graph.
	v1 := &Node{id: 1}
	v2 := &Node{id: 2}
	v3 := &Node{id: 3}
	g.AddNode(v1)
	g.AddNode(v2)
	g.AddNode(v3)

	// Add edges to the graph.
	g.SetEdge(g.NewEdge(v1, v2))
	g.SetEdge(g.NewEdge(v2, v3))

	// Visualize the graph.
	dotFile, err := dot.Marshal(g, "my-graph", "", "")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dotFile))
}
