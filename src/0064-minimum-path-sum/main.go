package main

// T: O(n*m))
// M: O(1)
// -- start --

func minPathSum(grid [][]int) int {
	colSize := len(grid)
	if colSize == 0 {
		return 0
	}
	rowSize := len(grid[0])

	// first row
	for row := 1; row < rowSize; row++ {
		grid[0][row] = grid[0][row-1] + grid[0][row]
	}

	// first col
	for col := 1; col < colSize; col++ {
		grid[col][0] = grid[col-1][0] + grid[col][0]
	}

	for col := 1; col < colSize; col++ {
		for row := 1; row < rowSize; row++ {
			grid[col][row] = Min(grid[col-1][row], grid[col][row-1]) + grid[col][row]
		}
	}

	return grid[colSize-1][rowSize-1]
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// -- end --
