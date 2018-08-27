package main

import (
	"fmt"
	"os"
)

func main() {
	maze := readMaze("maze/maze.in")
	//for _, row := range maze {
	//	for _, col := range row {
	//		fmt.Printf("%d ", col)
	//	}
	//	fmt.Println()
	//}
	steps := walk(maze, point{0, 0},
		point{len(maze) - 1, len(maze[0]) - 1})

	fmt.Println()
	for _, row := range steps {
		for _, col := range row {
			fmt.Printf("%2d ", col)
		}
		fmt.Println()
	}
}

type point struct {
	i, j int
}

// 上下左右四个方向
var dirs = [4]point{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false // 越界
	}

	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false // 越界
	}
	return grid[p.i][p.j], true
}

func walk(maze [][]int, start, end point) [][]int {
	// 保存步数 , 跟 maze结构一样，只是默认值全是0
	steps := make([][]int, len(maze))
	for i := range maze {
		steps[i] = make([]int, len(maze[i]))
	}

	Q := []point{start}

	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]

		if cur == end {
			break
		}

		for _, dir := range dirs {
			next := cur.add(dir)

			// 判断边界
			// maze at next is 0
			// and steps at next is 0
			// next != start
			val, ok := next.at(maze)
			if !ok || val == 1 {
				continue
			}

			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}

			if next == start {
				continue
			}

			curStep, _ := cur.at(steps)
			steps[next.i][next.j] = curStep + 1

			Q = append(Q, next)
		}
	}
	return steps
}

func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)
	//fmt.Println(row, col)

	// 行
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}
