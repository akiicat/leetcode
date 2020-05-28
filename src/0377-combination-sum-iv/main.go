package main

// T: O(n * t)
// M: O(t)
// -- start --

// T: O(n * t)
// M: O(t)
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1

	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if i-num >= 0 {
				dp[i] += dp[i-num]
			}
		}
	}

	return dp[target]
}

// T: O(n * t)
// M: O(t)
func combinationSum4Recursive(nums []int, target int) int {
	cache := make(map[int]int)
	return R(cache, nums, target)
}

func R(cache map[int]int, nums []int, target int) int {
	if v, ok := cache[target]; ok {
		return v
	}

	if target == 0 {
		return 1
	}
	if target < 0 {
		return 0
	}

	s := 0
	for _, v := range nums {
		s += R(cache, nums, target-v)
	}
	cache[target] = s
	return s
}

// -- end --
