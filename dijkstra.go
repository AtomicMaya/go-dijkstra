package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"sort"
)

type Vertex struct {
	Label    string
	Distance int32
}

type Edge struct {
	Start *Vertex
	End   *Vertex
	Value int32
}

func dijkstra(origin *Vertex, destination *Vertex, v []*Vertex, e []*Edge) []string {
	fmt.Println("Starting")
	origin.Distance = 0
	heap := VertexQueue{Elements: []*Vertex{origin}}
	visited := VertexQueue{Elements: []*Vertex{}}
	node, err := &Vertex{}, errors.New("")

	for heap.Size() > 0 {
		node, err = heap.Pop()
		fmt.Printf(`Node %v`, node)
		fmt.Println()
		if err != nil {
			fmt.Println(errors.New("no nodes in queue"))
			os.Exit(3)
		}
		visited.Append(node)
		filteredEdges := make([]*Edge, 0, len(e))
		for _, edge := range e {
			if edge.Start.Label == node.Label || edge.End.Label == node.Label {
				filteredEdges = append(filteredEdges, edge)
			}
		}

		sort.Slice(filteredEdges, func(i, j int) bool {
			return filteredEdges[i].Value > filteredEdges[j].Value
		})

		// If the vertex is a well
		if len(filteredEdges) == 1 &&
			(filteredEdges[0].Start != origin && filteredEdges[0].End != destination &&
				filteredEdges[0].Start != destination && filteredEdges[0].End != origin) {
			// Count the amount of times said node
			for _, edge := range e {
				if edge.Start == filteredEdges[0].Start || edge.End == filteredEdges[0].Start {
					filteredEdges[0].Start.Distance = math.MaxInt32
				} else if edge.End == filteredEdges[0].End || edge.Start == filteredEdges[0].End {
					filteredEdges[0].End.Distance = math.MaxInt32
				}
				filteredEdges[0].Value = math.MaxInt32
			}
		} else {
			// Iterate on all available edges.
			for _, edge := range filteredEdges {
				start, end := &Vertex{}, &Vertex{}
				if edge.Start.Label == node.Label {
					start = edge.Start
					end = edge.End
				} else {
					start = edge.End
					end = edge.Start
				}

				if end.Label != origin.Label && ((end.Distance == 0 && end.Distance < start.Distance+edge.Value) || (end.Distance > start.Distance+edge.Value)) {
					end.Distance = start.Distance + edge.Value
				}

				if node.Label == destination.Label {
					heap.DequeueWhere(func(v Vertex) bool { return true })
				} else if !visited.Contains(*end) && ((heap.Contains(*end) && heap.FilterWhere(func(v Vertex) bool { return v.Label == end.Label })[0].Distance > end.Distance) || !heap.Contains(*end)) {
					heap.Append(end)
				}
			}

		}

		sort.Slice(heap.Elements, func(i, j int) bool {
			return heap.Elements[i].Distance > heap.Elements[j].Distance
		})

		for _, vec := range v {
			fmt.Println(*vec)
		}

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

	vertices := []*Vertex{&a, &b, &c, &d, &f, &g, &h, &i}

	edges := []*Edge{
		{Start: &a, End: &b, Value: 1},
		{Start: &b, End: &c, Value: 3},
		{Start: &b, End: &d, Value: 2},
		{Start: &b, End: &f, Value: 5},
		{Start: &b, End: &g, Value: 2},
		{Start: &c, End: &d, Value: 3},
		{Start: &c, End: &h, Value: 7},
		{Start: &c, End: &g, Value: 5},
		{Start: &d, End: &i, Value: 4},
		{Start: &f, End: &g, Value: 2},
		{Start: &i, End: &h, Value: 8},
		{Start: &h, End: &j, Value: 9},
	}

	fmt.Println(dijkstra(&i, &f, vertices, edges))
}
