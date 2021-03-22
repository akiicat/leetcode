package main

// T: O(n)
// M: O(1)
// -- start --

// T: O(n)
// M: O(1)
func jump(nums []int) int {
	curMax := 0
	count := 0
	nextMax := 0

	for cur, num := range nums {
		if cur <= curMax {
			nextMax = Max(nextMax, cur+num)
		}

		if cur == len(nums)-1 {
			break
		}

		if cur == curMax {
			curMax = nextMax
			count++
		}
	}

	return count
}

// T: O(n ^ 2)
// M: O(n)
func jumpDP(nums []int) int {
	arr := make([]int, len(nums))

	for i, _ := range nums {
		if i == 0 {
			continue
		}

		min := 1 << 31
		for j := i - 1; j >= 0; j-- {
			if nums[j] >= i-j {
				min = Min(min, arr[j]+1)
			}
		}
		arr[i] = min
	}

	return arr[len(nums)-1]
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// -- end --
