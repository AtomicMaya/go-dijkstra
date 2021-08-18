package main

import (
	"testing"
)

func TestSize(t *testing.T) {
	queue := VertexQueue{Elements: []*Vertex{
		{Label: "A", Distance: 0},
	}}

	if queue.Size() != 1 {
		t.Fatalf(`Queue.Size() = %v, %v`, queue.Size(), 1)
	}
}

func TestShift(t *testing.T) {
	queue := VertexQueue{Elements: []*Vertex{
		{Label: "B", Distance: 0},
		{Label: "C", Distance: 0},
		{Label: "D", Distance: 0},
	}}

	testVertex, err := queue.Shift()

	if err != nil {
		t.Fatal("Shift operation failed")
	} else if testVertex.Label != "B" {
		t.Fatalf(`Queue.Shift() = %v, %v`, testVertex, Vertex{Label: "B", Distance: 0})
	}

	if queue.Size() != 2 || queue.Elements[0].Label != "C" {
		t.Fatalf(`Queue.Size() = %v, %v, First = %s, %s`, queue.Size(), 2, queue.Elements[0].Label, "C")
	}
}

func TestPop(t *testing.T) {
	queue := VertexQueue{Elements: []*Vertex{
		{Label: "B", Distance: 0},
		{Label: "C", Distance: 0},
		{Label: "D", Distance: 0},
	}}

	testVertex, err := queue.Pop()

	if err != nil {
		t.Fatal("Pop operation failed")
	} else if testVertex.Label != "D" {
		t.Fatalf(`Queue.Pop() = %v, %v`, testVertex, Vertex{Label: "D", Distance: 0})
	}

	if queue.Size() != 2 || queue.Elements[0].Label != "B" {
		t.Fatalf(`Queue.Size() = %v, %v, First = %s, %s`, queue.Size(), 2, queue.Elements[0].Label, "C")
	}
}

func TestAppend(t *testing.T) {
	queue := VertexQueue{Elements: []*Vertex{
		{Label: "B", Distance: 0},
		{Label: "C", Distance: 0},
		{Label: "D", Distance: 0},
	}}

	testVertex := Vertex{Label: "A", Distance: 0}

	queue.Append(&testVertex)

	if queue.Size() != 4 {
		t.Fatalf(`Queue.Size() = %v, %v`, queue.Size(), 4)
	}

	if queue.Elements[3].Label != "A" {
		t.Fatalf(`Queue.Append() => %v, %v`, queue.Elements[3], testVertex)
	}
}

func TestPrepend(t *testing.T) {
	queue := VertexQueue{Elements: []*Vertex{
		{Label: "B", Distance: 0},
		{Label: "C", Distance: 0},
		{Label: "D", Distance: 0},
	}}

	testVertex := Vertex{Label: "A", Distance: 0}

	queue.Prepend(&testVertex)

	if queue.Size() != 4 {
		t.Fatalf(`Queue.Size() = %v, %v`, queue.Size(), 4)
	}

	if queue.Elements[0].Label != "A" || queue.Elements[1].Label != "B" {
		t.Fatalf(`Queue.Prepend() => %v, %v`, queue.Elements[0], testVertex)
	}
}

func TestFirst(t *testing.T) {
	queue := VertexQueue{Elements: []*Vertex{
		{Label: "B", Distance: 0},
		{Label: "C", Distance: 0},
		{Label: "D", Distance: 0},
	}}

	testVertex, err := queue.First()

	if err != nil {
		t.Fatalf(`Queue.Size() = %v, %v`, queue.Size(), 3)
	} else if testVertex.Label != "B" {
		t.Fatalf(`Queue.First() = %v, %v`, testVertex, Vertex{Label: "B", Distance: 0})
	}
}

func TestLast(t *testing.T) {
	queue := VertexQueue{Elements: []*Vertex{
		{Label: "B", Distance: 0},
		{Label: "C", Distance: 0},
		{Label: "D", Distance: 0},
	}}

	testVertex, err := queue.Last()

	if err != nil {
		t.Fatalf(`Queue.Size() = %v, %v`, queue.Size(), 3)
	} else if testVertex.Label != "D" {
		t.Fatalf(`Queue.Last() = %v, %v`, testVertex, Vertex{Label: "D", Distance: 0})
	}
}

func TestDequeueWhere(t *testing.T) {
	queue := VertexQueue{Elements: []*Vertex{
		{Label: "B", Distance: 0},
		{Label: "C", Distance: 0},
		{Label: "D", Distance: 0},
	}}

	queue.DequeueWhere(func(v Vertex) bool {
		return v.Label != "C"
	})

	if queue.Size() != 1 {
		t.Fatalf(`Queue.Size() = %v, %v`, queue.Size(), 1)
	} else if queue.Elements[0].Label != "C" {
		t.Fatalf(`Queue[0] = %v, %v`, queue.Elements[0], Vertex{Label: "C", Distance: 0})
	}
}

func TestDequeueAll(t *testing.T) {
	queue := VertexQueue{Elements: []*Vertex{
		{Label: "B", Distance: 0},
		{Label: "C", Distance: 0},
		{Label: "D", Distance: 0},
	}}

	queue.DequeueWhere(func(v Vertex) bool {
		return true
	})

	if queue.Size() != 0 {
		t.Fatalf(`Queue.Size() = %v, %v`, queue.Size(), 0)
	}
}

func TestContains(t *testing.T) {
	queue := VertexQueue{Elements: []*Vertex{
		{Label: "B", Distance: 0},
		{Label: "C", Distance: 0},
		{Label: "D", Distance: 0},
	}}

	test1 := queue.Contains(Vertex{Label: "B", Distance: 0})
	test2 := queue.Contains(Vertex{Label: "A", Distance: 0})

	if !test1 {
		t.Fatalf(`Queue.Contains(Vertex{Label: "B", Distance: 0}) = %v, %v`, test1, true)
	}
	if test2 {
		t.Fatalf(`Queue.Contains(Vertex{Label: "A", Distance: 0}) = %v, %v`, test2, false)
	}
}

func TestFilterWhere(t *testing.T) {
	queue := VertexQueue{Elements: []*Vertex{
		{Label: "B", Distance: 0},
		{Label: "C", Distance: 0},
		{Label: "D", Distance: 0},
	}}

	filtered := queue.FilterWhere(func(v Vertex) bool { return v.Label != "C" })
	if len(filtered) != 2 {
		t.Fatalf(`Filtered Length = %v, %v`, len(filtered), 2)
	}
}
