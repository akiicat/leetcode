package main

// T: O(n!)
// M: O(n)
// -- start --

func permute(nums []int) [][]int {
	res := [][]int{}
	R(&res, 0, nums)
	return res
}

func R(res *[][]int, cur int, nums []int) {
	if cur == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*res = append(*res, tmp)
		return
	}

	for i := cur; i < len(nums); i++ {
		nums[cur], nums[i] = nums[i], nums[cur]
		R(res, cur+1, nums)
		nums[cur], nums[i] = nums[i], nums[cur]
	}
}

// -- end --
