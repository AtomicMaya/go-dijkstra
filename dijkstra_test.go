package main

import (
	"reflect"
	"testing"
)

func getTestData() ([]*Vertex, []*Edge) {
	a, b, c, d, f, g, h, i, j := Vertex{Label: "A", Distance: 0},
		Vertex{Label: "B", Distance: 0},
		Vertex{Label: "C", Distance: 0},
		Vertex{Label: "D", Distance: 0},
		Vertex{Label: "F", Distance: 0},
		Vertex{Label: "G", Distance: 0},
		Vertex{Label: "H", Distance: 0},
		Vertex{Label: "I", Distance: 0},
		Vertex{Label: "J", Distance: 0}

	vertices := []*Vertex{&a, &b, &c, &d, &f, &g, &h, &i, &j}

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

	return vertices, edges
}

func TestPath(t *testing.T) {
	vertices, edges := getTestData()

	result := dijkstra(vertices[7], vertices[4], vertices, edges)

	if !reflect.DeepEqual(result, []string{"I", "D", "B", "G", "F"}) {
		t.Fatalf(`PATH = %v, %v`, result, []string{"I", "D", "B", "G", "F"})
	}
}

func TestDestinationIsOrigin(t *testing.T) {
	vertices, edges := getTestData()

	result := dijkstra(vertices[7], vertices[7], vertices, edges)

	if !reflect.DeepEqual(result, []string{vertices[7].Label}) {
		t.Fatalf(`PATH = %v, %v`, result, []string{vertices[7].Label})
	}

	result = dijkstra(vertices[2], vertices[2], vertices, edges)

	if !reflect.DeepEqual(result, []string{vertices[2].Label}) {
		t.Fatalf(`PATH = %v, %v`, result, []string{vertices[2].Label})
	}
}

func TestOriginIsWell(t *testing.T) {
	vertices, edges := getTestData()

	result := dijkstra(vertices[0], vertices[6], vertices, edges)

	if !reflect.DeepEqual(result, []string{"A", "B", "C", "H"}) {
		t.Fatalf(`PATH = %v, %v`, result, []string{"A", "B", "C", "H"})
	}
}

func TestDestinationIsWell(t *testing.T) {
	vertices, edges := getTestData()

	result := dijkstra(vertices[5], vertices[0], vertices, edges)

	if !reflect.DeepEqual(result, []string{"G", "B", "A"}) {
		t.Fatalf(`PATH = %v, %v`, result, []string{"G", "B", "A"})
	}
}
