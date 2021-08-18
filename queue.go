package main

import (
	"errors"
)

type VertexQueue struct {
	Elements []*Vertex
}

func (queue *VertexQueue) Size() int {
	return len(queue.Elements)
}

/** Removes the first element in the queue and shifts the entire queue forward. */
func (queue *VertexQueue) Shift() (*Vertex, error) {
	if queue.Size() < 0 {
		return &Vertex{}, errors.New("the queue is empty")
	}
	element := queue.Elements[0]
	queue.Elements = queue.Elements[1:]
	return element, nil
}

/** Removes the last element of the queue and returns it. */
func (queue *VertexQueue) Pop() (*Vertex, error) {
	if queue.Size() < 0 {
		return &Vertex{}, errors.New("the queue is empty")
	}
	element := queue.Elements[queue.Size()-1]
	queue.Elements = queue.Elements[:queue.Size()-1]
	return element, nil
}

// Adds an element
func (queue *VertexQueue) Append(elements ...*Vertex) {
	queue.Elements = append(queue.Elements, elements...)
}

func (queue *VertexQueue) Prepend(elements ...*Vertex) {
	values := queue.Elements[:]
	queue.Elements = append([]*Vertex{}, elements...)
	queue.Elements = append(queue.Elements, values...)
}

func (queue *VertexQueue) First() (*Vertex, error) {
	if queue.Size() == 0 {
		return &Vertex{}, errors.New("the queue is empty")
	}
	return queue.Elements[0], nil
}

func (queue *VertexQueue) Last() (*Vertex, error) {
	if queue.Size() == 0 {
		return &Vertex{}, errors.New("the queue is empty")
	}
	return queue.Elements[queue.Size()-1], nil
}

func (queue *VertexQueue) DequeueWhere(function func(Vertex) bool) {
	elementsKept := []*Vertex{}
	for _, vertex := range queue.Elements {
		if !function(*vertex) {
			elementsKept = append(elementsKept, vertex)
		}
	}

	queue.Elements = elementsKept[:]
}

func (queue *VertexQueue) FilterWhere(function func(Vertex) bool) []*Vertex {
	elementsFiltered := []*Vertex{}

	for _, vertex := range queue.Elements {
		if function(*vertex) {
			elementsFiltered = append(elementsFiltered, vertex)
		}
	}

	return elementsFiltered
}

func (queue *VertexQueue) Contains(vertex Vertex) bool {
	for _, v := range queue.Elements {
		if vertex.Label == v.Label && vertex.Distance == v.Distance {
			return true
		}
	}
	return false
}
