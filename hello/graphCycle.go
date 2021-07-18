package hello

import "fmt"

type Graph struct {
	nodeByID map[int]*GraphVertex
}

type GraphVertex struct {
	ID    int
	edges []*GraphVertex
}

func (g *Graph) Add(ID int) *GraphVertex {
	if node, ok := g.nodeByID[ID]; ok {
		return node
	}
	g.nodeByID[ID] = &GraphVertex{ID: ID}
	return g.nodeByID[ID]
}

func New() *Graph {
	return &Graph{map[int]*GraphVertex{}}
}

func (v *GraphVertex) AddEdges(vertices ...*GraphVertex) {
	v.edges = append(v.edges, vertices...)
}

func (g *Graph) HasCycle() bool {
	seen := map[int]struct{}{}
	for _, node := range g.nodeByID {
		seenInCycle := map[int]struct{}{}
		queue := []*GraphVertex{node}
		for len(queue) > 0 {
			curr := queue[0]
			queue[0] = nil
			queue = queue[1:]

			if _, ok := seenInCycle[curr.ID]; ok {
				return true
			}
			if _, ok := seen[curr.ID]; ok {
				continue
			}

			seenInCycle[curr.ID] = struct{}{}
			seen[curr.ID] = struct{}{}
			queue = append(queue, curr.edges...)
		}
	}
	return false
}

func IsValidSchedule(sched [][]int) bool {
	g := New()
	for _, v := range sched {
		edges := []*GraphVertex{}
		for _, ID := range v[1:] {
			edges = append(edges, g.Add(ID))
		}
		g.Add(v[0]).AddEdges(edges...)
	}
	return !g.HasCycle()
}

func TestCourseSchedule() {
	fmt.Println(IsValidSchedule([][]int{{1, 0}, {0, 1}}))
	fmt.Println(IsValidSchedule([][]int{{1, 0}}))
	fmt.Println(IsValidSchedule([][]int{{1, 0}, {0, 2}}))
}
