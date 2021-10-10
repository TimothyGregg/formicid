package graph

import (
	"fmt"
	"math"

	"github.com/fogleman/delaunay"
)

type Vertex struct {
	X int `json:"x"`
	Y int `json:"y`
}

func (v *Vertex) Position() (int, int) {
	return v.X, v.Y
}

func (v Vertex) String() string {
	return "(" + fmt.Sprint(v.X) + ", " + fmt.Sprint(v.Y) + ")"
}

func (v1 *Vertex) Same_As(v2 *Vertex) bool {
	return v1.X == v2.X && v1.Y == v2.Y
}

type Edge struct {
	v1, v2 *Vertex
	length float64
}

func (e *Edge) Length() float64 {
	return e.length
}

func NewEdge(v1, v2 *Vertex) *Edge {
	e := &Edge{v1: v1, v2: v2}
	e.length = float64(math.Sqrt(math.Pow(float64(v2.X-v1.X), 2) + math.Pow(float64(v2.Y-v1.Y), 2)))
	return e
}

func (e *Edge) Vertices() (*Vertex, *Vertex) {
	return e.v1, e.v2
}

func (e Edge) String() string {
	return fmt.Sprint(e.v1) + " to " + fmt.Sprint(e.v2)
}

func (e1 *Edge) same_as(e2 *Edge) bool {
	return (e1.v1 == e2.v1 && e1.v2 == e2.v2) || (e1.v1 == e2.v2 && e1.v2 == e2.v1)
}

type Graph struct {
	Vertices  []*Vertex
	Edges     []*Edge
	Adjacency map[*Vertex][]*connection // We using an Adjaceny List boys. |E|/|V|^2 is typically > 1/64, at least in the graphs I like seeing it make
}

type connection struct {
	edge   *Edge
	vertex *Vertex
}

func New_Graph() *Graph {
	g := &Graph{}
	g.Adjacency = make(map[*Vertex][]*connection)
	return g
}

func (g Graph) String() string {
	outstr := ""
	for i, vertex := range g.Vertices {
		outstr = outstr + "[" + fmt.Sprint(i) + "]: " + vertex.String() + "\n"
	}
	for _, edge := range g.Edges {
		outstr = outstr + edge.String() + "; "
	}
	return outstr
}

func (g *Graph) has(v *Vertex) bool {
	for _, v_test := range g.Vertices {
		if v == v_test {
			return true
		}
	}
	return false
}

func (g *Graph) Add_Vertex(x, y int) (*Vertex, error) {
	v := &Vertex{X: x, Y: y}
	for _, v_test := range g.Vertices {
		if v.Same_As(v_test) {
			return v_test, &VertexAlreadyExistsError{vertex: v_test}
		}
	}
	g.Vertices = append(g.Vertices, v)
	return v, nil
}

func (g *Graph) Add_Edge(v1 *Vertex, v2 *Vertex) (*Edge, error) {
	if !g.has(v1) {
		return nil, &MissingVertexError{vertex: v1}
	}
	if !g.has(v2) {
		return nil, &MissingVertexError{vertex: v2}
	}
	e := NewEdge(v1, v2)
	for _, c_test := range g.Adjacency[v1] {
		if e.same_as(c_test.edge) {
			return c_test.edge, &EdgeAlreadyExistsError{edge: c_test.edge}
		}
	}
	g.Edges = append(g.Edges, e)
	g.Adjacency[v1] = append(g.Adjacency[v1], &connection{vertex: v2, edge: e})
	g.Adjacency[v2] = append(g.Adjacency[v2], &connection{vertex: v1, edge: e})
	return e, nil
}

func (g *Graph) Remove_Edge(e *Edge) error {
	for it, edge := range g.Edges {
		if e.same_as(edge) {
			// Remove from graph Edge list
			g.Edges = append(g.Edges[:it], g.Edges[it+1:]...)
			// Remove from the two Adjacency slices
			for it, conn := range g.Adjacency[e.v1] {
				if e.same_as(conn.edge) {
					g.Adjacency[e.v1] = append(g.Adjacency[e.v1][:it], g.Adjacency[e.v1][it+1:]...)
					break
				}
			}
			for it, conn := range g.Adjacency[e.v2] {
				if e.same_as(conn.edge) {
					g.Adjacency[e.v2] = append(g.Adjacency[e.v2][:it], g.Adjacency[e.v2][it+1:]...)
					break
				}
			}
			return nil
		}
	}
	return &EdgeNotFoundError{edge: e}
}

func (g *Graph) Delaunay_Triangulate() (*delaunay.Triangulation, error) {
	var points []delaunay.Point
	for _, v := range g.Vertices {
		points = append(points, delaunay.Point{X: float64(v.X), Y: float64(v.Y)})
	}
	return delaunay.Triangulate(points)
}

func (g *Graph) Connect_Delaunay() error { // https://mapbox.github.io/delaunator/
	triangulation, err := g.Delaunay_Triangulate()
	for it := 0; it < len(triangulation.Triangles)/3; it++ {
		for jt := 0; jt < 3; jt++ {
			g.Add_Edge(g.Vertices[triangulation.Triangles[3*it+jt]], g.Vertices[triangulation.Triangles[3*it+(jt+1)%3]])
		}
	}
	return err
}
