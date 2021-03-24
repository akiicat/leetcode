package main

// T: O(min(m,n)) // min(m, n)
// M: O(1)
// -- start --

func uniquePaths(m int, n int) int {
	// let m >= n
	if m < n {
		m, n = n, m
	}

	res := 1

	for i := 1; i <= (n - 1); i++ {
		res = res * (m + i - 1) / i
	}

	return res
}

// -- end --
