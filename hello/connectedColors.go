package hello

import "fmt"

type Grid [][]Color
type Color int
type Point struct{ x, y int }
type ConnectedColorState struct {
	Visited map[Point]struct{}
}

func (state ConnectedColorState) hasVisited(point Point) bool {
	_, ok := state.Visited[point]
	return ok
}

func (state ConnectedColorState) visit(point Point) {
	state.Visited[point] = struct{}{}
}

func (point Point) getNeighbors(grid Grid) []Point {
	x, y := point.x, point.y
	validNeighbors := make([]Point, 0)
	possibleNeighbors := []Point{
		{x - 1, y}, {x + 1, y}, {x, y + 1}, {x, y - 1},
	}
	for _, n := range possibleNeighbors {
		if n.validForGrid(grid) {
			validNeighbors = append(validNeighbors, n)
		}
	}
	return validNeighbors
}

func (point Point) validForGrid(grid Grid) bool {
	return point.x >= 0 && point.y >= 0 && point.x < grid.Width() && point.y < grid.Height()
}

func intMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (grid Grid) Height() int {
	return len(grid)
}

func (grid Grid) Width() int {
	return len(grid[0])
}

func (grid Grid) ColorOf(point Point) Color {
	return grid[point.y][point.x]
}

func findConnectedColors(grid Grid, state ConnectedColorState, start Point) int {
	if state.hasVisited(start) {
		return 0
	}
	state.visit(start)
	color := grid.ColorOf(start)
	queue := []Point{start}
	length := 1
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		for _, neighbor := range p.getNeighbors(grid) {
			if grid.ColorOf(neighbor) == color && !state.hasVisited(neighbor) {
				length += 1
				state.visit(neighbor)
				queue = append(queue, neighbor)
			}
		}
	}
	return length
}

func (grid Grid) MaxConnectedColors() int {
	result := 0
	state := ConnectedColorState{map[Point]struct{}{}}
	for y, row := range grid {
		for x := range row {
			result = intMax(result, findConnectedColors(grid, state, Point{x, y}))
		}
	}
	return result
}

func TestGridMaxConnectedColors() {
	fmt.Println(Grid{
		{1, 1, 1},
		{2, 3, 1},
		{2, 1, 1},
	}.MaxConnectedColors())
}
