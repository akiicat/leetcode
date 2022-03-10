package main

// T: O(log(n))
// M: O(1)
// -- start --

func searchRange(nums []int, target int) []int {

	l, r := 0, len(nums)
	for l < r {
		m := (l + r) / 2

		if nums[m] >= target {
			r = m
		} else {
			l = m + 1
		}
	}

	lo := l

	l, r = 0, len(nums)
	for l < r {
		m := (l + r) / 2

		if nums[m] > target {
			r = m
		} else {
			l = m + 1
		}
	}

	hi := l

	if lo == hi {
		return []int{-1, -1}
	}

	return []int{lo, hi - 1}
}

// -- end --
