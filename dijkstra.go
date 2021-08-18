package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"sort"
)

/** Defined by a Label and a distance (default 0), the minimum distance from the starting node. */
type Vertex struct {
	Label    string
	Distance int32
}

/** Defined by two vertices and a value representing the edge's length. */
type Edge struct {
	Start *Vertex
	End   *Vertex
	Value int32
}

/**
For a graph defined by vertices v and edges e, returns the cheapest path between origin and destination.
*/
func dijkstra(origin *Vertex, destination *Vertex, v []*Vertex, e []*Edge) []string {
	// Forcing our origin to start at distance 0 (in case of the algorithm running multiple times on the same data)
	origin.Distance = 0

	// Initialisation of various Queues
	queue := VertexQueue{Elements: []*Vertex{origin}}
	visited := VertexQueue{Elements: []*Vertex{}}
	node, err := &Vertex{}, errors.New("")

	// Iterate over all of the elements of the queue until there are no more vertices (max O(len(v)))
	for queue.Size() > 0 {
		// Always work on what is closest to the current vertex in the queue.
		node, err = queue.Pop()
		if err != nil {
			fmt.Println(errors.New("no nodes in queue"))
			os.Exit(3)
		}

		// Avoid repetitions
		visited.Append(node)

		// Filtering out the edges that are linked to the current node
		filteredEdges := make([]*Edge, 0, len(e))
		for _, edge := range e {
			if edge.Start.Label == node.Label || edge.End.Label == node.Label {
				filteredEdges = append(filteredEdges, edge)
			}
		}

		// Sorting the edges by distance
		sort.Slice(filteredEdges, func(i, j int) bool {
			return filteredEdges[i].Value > filteredEdges[j].Value
		})

		// If the vertex is a well
		if len(filteredEdges) == 1 &&
			(filteredEdges[0].Start != origin && filteredEdges[0].End != destination &&
				filteredEdges[0].Start != destination && filteredEdges[0].End != origin) {
			// Set the node to be unreachable by the backtrace
			for _, edge := range e {
				filteredEdges[0].Value = math.MaxInt32
				if edge.Start == filteredEdges[0].Start || edge.End == filteredEdges[0].Start {
					filteredEdges[0].Start.Distance = math.MaxInt32
					break
				} else if edge.End == filteredEdges[0].End || edge.Start == filteredEdges[0].End {
					filteredEdges[0].End.Distance = math.MaxInt32
					break
				}
			}
		} else {
			// Iterate on all available edges.
			for _, edge := range filteredEdges {
				// Determine the actual direction of the edge.
				start, end := &Vertex{}, &Vertex{}
				if edge.Start.Label == node.Label {
					start = edge.Start
					end = edge.End
				} else {
					start = edge.End
					end = edge.Start
				}

				// If end distance not yet set.
				if end.Label != origin.Label && ((end.Distance == 0 && end.Distance < start.Distance+edge.Value) || (end.Distance > start.Distance+edge.Value)) {
					end.Distance = start.Distance + edge.Value
				}

				// If arrival at destination, empty queue.
				if node.Label == destination.Label {
					queue.DequeueWhere(func(v Vertex) bool { return true })
					// Otherwise append all further nodes to the queue.
				} else if !visited.Contains(*end) && ((queue.Contains(*end) && queue.FilterWhere(func(v Vertex) bool { return v.Label == end.Label })[0].Distance > end.Distance) || !queue.Contains(*end)) {
					queue.Append(end)
				}
			}
		}

		sort.Slice(queue.Elements, func(i, j int) bool {
			return queue.Elements[i].Distance > queue.Elements[j].Distance
		})
	}

	// Path of labels for the trace.
	path := []*Vertex{destination}
	queue = VertexQueue{Elements: []*Vertex{destination}}
	visited = VertexQueue{Elements: []*Vertex{}}
	node, err = &Vertex{}, errors.New("")

	// Backtrace
	for queue.Size() > 0 {
		node, err = queue.Pop()
		if err != nil {
			fmt.Println(errors.New("no nodes in queue"))
			os.Exit(3)
		}
		visited.Append(node)
		filteredEdges := make([]*Edge, 0, len(e))
		for _, edge := range e {
			if (edge.Start.Label == node.Label || edge.End.Label == node.Label) && edge.Value != math.MaxInt32 {
				filteredEdges = append(filteredEdges, edge)
			}
		}

		sort.Slice(filteredEdges, func(i, j int) bool {
			return filteredEdges[i].Value > filteredEdges[j].Value
		})

		for _, edge := range filteredEdges {
			start, end := &Vertex{}, &Vertex{}
			if edge.Start.Label == node.Label {
				start = edge.Start
				end = edge.End
			} else {
				start = edge.End
				end = edge.Start
			}

			if node.Label == origin.Label {
				queue.DequeueWhere(func(_ Vertex) bool { return true })
			} else if !visited.Contains(*end) && !queue.Contains(*end) && start.Distance-end.Distance == edge.Value {
				// If the node hasn't been visited and isn't planned for visit yet, and that the edge value corresponds to the delta of the distance from the start.
				queue.Append(end)
				path = append(path, end)
			}
		}

		sort.Slice(queue.Elements, func(i, j int) bool {
			return queue.Elements[i].Distance > queue.Elements[j].Distance
		})
	}

	// Reverse the path to obtain the path in the right direction.
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	// Format to string
	result := []string{}
	for _, vertex := range path {
		result = append(result, vertex.Label)
	}
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
