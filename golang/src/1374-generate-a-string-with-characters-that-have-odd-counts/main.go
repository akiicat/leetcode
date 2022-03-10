package main

// T: O(n)
// M: O(1)
// -- start --

func generateTheString(n int) string {
	if n == 0 {
		return ""
	}

	// even
	if n%2 == 0 {
		return generateTheString(n-1) + "b"
	}

	// odd
	res := ""

	for n > 0 {
		res += "a"
		n--
	}

	return res
}

// -- end --
