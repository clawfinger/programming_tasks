package main

import "fmt"

func bfs(field [][]byte, marked [][]int, i int, j int) {
	if i < 0 || j < 0 || i >= len(field) ||
		j >= len(field[0]) || field[i][j] == byte('0') || marked[i][j] == 1 {
		return
	}
	if field[i][j] == '1' {
		marked[i][j] = 1
	}
	bfs(field, marked, i, j-1)
	bfs(field, marked, i, j+1)
	bfs(field, marked, i-1, j)
	bfs(field, marked, i+1, j)
}

func numIslands(grid [][]byte) int {
	marked := make([][]int, 0)
	for i := 0; i < len(grid); i++ {
		marked = append(marked, []int{})
		for j := 0; j < len(grid[i]); j++ {
			marked[i] = append(marked[i], 0)
		}
	}
	counter := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if marked[i][j] == 1 {
				continue
			}
			if grid[i][j] == '1' {
				counter++
			}
			bfs(grid, marked, i, j)
		}
	}
	return counter
}

func main() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	fmt.Println(numIslands(grid))
}
