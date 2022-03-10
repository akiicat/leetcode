package main

// T: O(n ^ 2) // n is the size of matrix
// M: O(1)
// -- start --

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	dirs := [][]int{
		[]int{0, 1},
		[]int{1, 0},
		[]int{0, -1},
		[]int{-1, 0},
	}

	bound := [][]int{
		[]int{1, 0, 0, 0},
		[]int{0, -1, 0, 0},
		[]int{0, 0, -1, 0},
		[]int{0, 0, 0, 1},
	}

	l, r, t, b := 0, len(matrix[0]), 0, len(matrix)

	d := 0

	x, y := 0, 0

	rtn := []int{}

	for l < r && t < b {
		nx, ny := x, y
		for t <= nx && nx < b && l <= ny && ny < r {
			rtn = append(rtn, matrix[nx][ny])
			x, y, nx, ny = nx, ny, nx+dirs[d][0], ny+dirs[d][1]
		}

		t, r, b, l = t+bound[d][0], r+bound[d][1], b+bound[d][2], l+bound[d][3]
		d = (d + 1) % 4

		x, y = x+dirs[d][0], y+dirs[d][1]
	}

	return rtn
}

// -- end --
