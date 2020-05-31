package main

import (
	"sort"
)

// T: O(n * n!)
// M: O(n)
// -- start --

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	R(&res, nums, 0)
	return res
}

func R(res *[][]int, nums []int, n int) {
	if n == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*res = append(*res, tmp)
		return
	}

	max := nums[n]
	for i := n; i < len(nums); i++ {
		if i == n {
			R(res, nums, n+1)
			continue
		}

		if nums[i] > max {
			max = nums[i]

			arr := []int{}
			arr = append(arr, nums[:n]...)
			arr = append(arr, nums[i])
			arr = append(arr, nums[n:i]...)
			arr = append(arr, nums[i+1:]...)

			R(res, arr, n+1)
		}
	}
}

// -- end --
