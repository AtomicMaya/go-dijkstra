package main

import (
	"fmt"
)

type Vertex struct {
	Label    string
	Distance int32
}

type Edge struct {
	Start Vertex
	End   Vertex
	Value int32
}

type SortByEdgeValue []Edge

func (a SortByEdgeValue) Len() int           { return len(a) }
func (a SortByEdgeValue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByEdgeValue) Less(i, j int) bool { return a[i].Value < a[j].Value }

func dijkstra(origin Vertex, destination Vertex, v []Vertex, e []Edge) []string {
	fmt.Println("Starting")
	origin.Distance = 0
	heap := make(chan Vertex)
	heap <- origin
	visited := make([]Vertex, len(v))
	node := Vertex{Label: "", Distance: 0}

	for len(heap) != 0 {
		node = <-heap
		visited = append(visited, node)
		filteredEdges := make([]Edge, 0, len(e))
		for _, edge := range e {
			if edge.Start.Label == node.Label || edge.End.Label == node.Label {
				filteredEdges = append(filteredEdges, edge)
			}
		}
		filteredEdges = SortByEdgeValue(filteredEdges)
		fmt.Println(filteredEdges)
	}
	result := make([]string, len(e))
	return result
}

func main() {
	a, b, c, d, f, g, h, i, j := Vertex{Label: "A", Distance: 0},
		Vertex{Label: "B", Distance: 0},
		Vertex{Label: "C", Distance: 0},
		Vertex{Label: "D", Distance: 0},
		Vertex{Label: "F", Distance: 0},
		Vertex{Label: "G", Distance: 0},
		Vertex{Label: "H", Distance: 0},
		Vertex{Label: "I", Distance: 0},
		Vertex{Label: "J", Distance: 0}

	vertices := []Vertex{a, b, c, d, f, g, h, i}

	edges := []Edge{
		Edge{Start: a, End: b, Value: 1},
		Edge{Start: b, End: c, Value: 3},
		Edge{Start: b, End: d, Value: 2},
		Edge{Start: b, End: f, Value: 5},
		Edge{Start: b, End: g, Value: 2},
		Edge{Start: c, End: d, Value: 3},
		Edge{Start: c, End: h, Value: 7},
		Edge{Start: c, End: g, Value: 5},
		Edge{Start: d, End: i, Value: 4},
		Edge{Start: f, End: g, Value: 2},
		Edge{Start: i, End: h, Value: 8},
		Edge{Start: h, End: j, Value: 9},
	}

	fmt.Println(dijkstra(i, f, vertices, edges))
}
