package main

import (
	"fmt"
	"os"
)

type point struct {
	i, j int
}

func (p point) next() []point {
	var newp = []point{
		{p.i, p.j - 1}, //left
		{p.i + 1, p.j}, //down
		{p.i, p.j + 1}, //right
		{p.i - 1, p.j}, //up
	}
	//fmt.Println(newp)
	return newp
}

func (p point) value(maze [][]int) int {
	if p.i < 0 || p.j < 0 {
		return 1
	}
	if p.i == len(maze) || p.j == len(maze[0]) {
		return 1
	}

	return maze[p.i][p.j]
}

func readmaze(s string) (maze, steps [][]int, row int, col int) {
	file, err := os.Open(s)
	if err != nil {
		panic(err)
	}

	fmt.Fscanf(file, "%d %d", &row, &col)

	maze = make([][]int, row)
	steps = make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		steps[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	//return maze and duplicate a blank Steps slice the same structure
	for _, i := range maze {
		for _, val := range i {
			print(val)

		}
		println()
	}
	return maze, steps, row, col

}

func walk(maze [][]int, steps [][]int, start point, end point) [][]int {
	pos := []point{start} //pos means position waiting to walk

	for len(pos) > 0 {
		cur := pos[0] // cur means currenty position
		pos = pos[1:]
		if cur == end {
			break
		}
		for _, next := range cur.next() {
			/*fmt.Print(next.i, next.j, next.value(maze))
			fmt.Println()*/
			if next == start || next.value(maze) == 1 {
				continue
			}

			if next.value(maze) == 0 && next.value(steps) == 0 {
				pos = append(pos, next)
				steps[next.i][next.j] = cur.value(steps) + 1
			}
		}
	}
	return steps
}

func main() {

	maze, steps, row, col := readmaze("src/learnGo/maze/maze.in")
	var start, end = point{0, 0}, point{row - 1, col - 1}
	steps = walk(maze, steps, start, end)
	fmt.Println()

	for _, i := range steps {
		for _, val := range i {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}
}
