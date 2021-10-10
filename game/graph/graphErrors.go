package graph

import "fmt"

type VertexAlreadyExistsError struct {
	vertex *Vertex
}

func (v *VertexAlreadyExistsError) Error() string {
	return fmt.Sprintf("Vertex already exists: %v", v.vertex)
}

type MissingVertexError struct {
	vertex *Vertex
}

func (m *MissingVertexError) Error() string {
	return fmt.Sprintf("Vertex does not exist in the graph: %v", m.vertex)
}

type EdgeAlreadyExistsError struct {
	edge *Edge
}

func (e *EdgeAlreadyExistsError) Error() string {
	return fmt.Sprintf("Edge already exists: %v", e.edge)
}

type EdgeNotFoundError struct {
	edge *Edge
}

func (e *EdgeNotFoundError) Error() string {
	return fmt.Sprintf("Edge not found: %v", e.edge)
}