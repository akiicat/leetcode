package main

// T: O(log(n))
// M: O(1)
// -- start --

// T: O(log(n))
// M: O(1)
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	if len(nums) > 0 && nums[l] == target {
		return l
	}

	for l < r {
		if nums[l] == target {
			return l
		}
		if nums[r] == target {
			return r
		}

		m := (l + r) / 2

		if nums[m] == target {
			return m
		}

		if nums[l] < nums[m] {
			if nums[l] < target && target < nums[m] {
				r = m - 1
			} else {
				l = m + 1
			}
		} else {
			if nums[m] < target && target < nums[r] {
				l = m + 1
			} else {
				r = m - 1
			}
		}
	}

	return -1
}

// T: O(n)
// M: O(1)
func searchLinear(nums []int, target int) int {
	for i, v := range nums {
		if v == target {
			return i
		}
	}

	return -1
}

// -- end --
