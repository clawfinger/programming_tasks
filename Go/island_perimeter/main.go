package main

func countLandNeighbours(grid [][]int, i int, j int) int {
	res := 0
	if j != 0 {
		res += grid[i][j-1]
	}
	if j != len(grid[i])-1 {
		res += grid[i][j+1]
	}
	if i != 0 {
		res += grid[i-1][j]

	}
	if i != len(grid)-1 {
		res += grid[i+1][j]

	}
	return res
}

func islandPerimeter(grid [][]int) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				res += 4 - countLandNeighbours(grid, i, j)

			}
		}
	}
	return res
}
