package main

// T: O(n^2)
// M: O(1)
// -- start --

func generateMatrix(n int) [][]int {
	res := make([][]int, n)

	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	dirs := [4][2]int{
		// col idx, row idx
		[2]int{0, 1},
		[2]int{1, 0},
		[2]int{0, -1},
		[2]int{-1, 0},
	}
	dir := 0

	col, row := 0, -1

	for i := 1; i <= n*n; i++ {
		nextCol, nextRow := col+dirs[dir][0], row+dirs[dir][1]
		for !(0 <= nextRow && nextRow < n && 0 <= nextCol && nextCol < n && res[nextCol][nextRow] == 0) {
			dir = (dir + 1) % 4
			nextCol, nextRow = col+dirs[dir][0], row+dirs[dir][1]
		}

		row, col = nextRow, nextCol

		res[col][row] = i
	}

	return res
}

// -- end --
