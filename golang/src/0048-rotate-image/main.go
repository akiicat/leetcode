package main

// T: O(n ^ 2) // n is the size of matrix
// M: O(1)
// -- start --

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2+1; i++ {
		for j := i; j < n-i-1; j++ {
			matrix[i][j], matrix[n-1-j][i], matrix[n-1-i][n-1-j], matrix[j][n-1-i] = matrix[n-1-j][i], matrix[n-1-i][n-1-j], matrix[j][n-1-i], matrix[i][j]
		}
	}
}

// -- end --
