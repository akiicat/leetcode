package main

// T: O(n)
// M: O(1)
// -- start --

// Greedy
// T: O(n)
// M: O(1)
func canJump(nums []int) bool {

	lastPos := len(nums) - 1

	for i := len(nums) - 1; i >= 0; i-- {
		if i+nums[i] >= lastPos {
			lastPos = i
		}
	}

	return lastPos == 0
}

// Dynamic Programming Top-Down
// T: O(n ^ 2)
// M: O(1)
func canJumpDynamicProgramming(nums []int) bool {
	m := make([]bool, len(nums))
	m[len(nums)-1] = true

	for i := len(nums) - 2; i >= 0; i-- {
		can := false

		for j := i + 1; j < len(nums) && j < i+nums[i]+1; j++ {
			can = can || m[j]
		}

		m[i] = can
	}

	return m[0]
}

// -- end --
