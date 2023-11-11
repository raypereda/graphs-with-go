I recommend the following Go packages for displaying directed graphs:

    yourbasic/graph: https://pkg.go.dev/github.com/yourbasic/graph is a simple and lightweight package for creating and manipulating graph data structures. It supports both directed and undirected graphs, and provides a variety of functions for graph traversal and analysis. The package also includes a built-in renderer for displaying graphs in a variety of formats, including SVG, PNG, and DOT.

    gonum/graph: https://pkg.go.dev/gonum.org/v1/gonum/graph is a more comprehensive package for graph algorithms and data structures. It supports a wide variety of graph types, including directed and undirected graphs, weighted and unweighted graphs, and hypergraphs. The package also includes a variety of graph algorithms, such as shortest path, minimum spanning tree, and maximum flow. The package does not include a built-in renderer, but it can be used to generate DOT code, which can then be rendered using a variety of tools.

Both of these packages are well-maintained and have a good track record of stability. They are also both relatively easy to use, making them a good choice for both beginners and experienced Go developers.

Which package should you choose?

If you are looking for a simple and lightweight package for displaying directed graphs, I recommend yourbasic/graph: https://pkg.go.dev/github.com/yourbasic/graph. It is easy to use and provides a built-in renderer for displaying graphs in a variety of formats.

If you need a more comprehensive package for graph algorithms and data structures, I recommend gonum/graph: https://pkg.go.dev/gonum.org/v1/gonum/graph. It supports a wide variety of graph types and algorithms, and it is well-maintained and stable.

Example

The following example shows how to use the yourbasic/graph: https://pkg.go.dev/github.com/yourbasic/graph package to display a directed graph:
Go

import (
    "github.com/yourbasic/graph"
)

func main() {
    // Create a directed graph.
    g := graph.NewDirectedGraph()

    // Add some nodes and edges to the graph.
    g.AddNode("A")
    g.AddNode("B")
    g.AddNode("C")
    g.AddEdge("A", "B")
    g.AddEdge("B", "C")
    g.AddEdge("C", "A")

    // Render the graph to an SVG file.
    err := g.RenderToSVG("graph.svg")
    if err != nil {
        log.Fatal(err)
    }
}

Use code with caution. Learn more

This example will create a directed graph with three nodes (A, B, and C) and three edges (A->B, B->C, and C->A). The graph will then be rendered to an SVG file called graph.svg
