package main

// T: O(n)
// M: O(1)
// -- start --

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}

	l := len(nums) - 2

	for l >= 0 {
		if nums[l] < nums[l+1] {

			min := l + 1

			for r := l + 1; r < len(nums); r++ {
				if nums[r] > nums[l] && nums[r] <= nums[min] {
					min = r
				}
			}

			nums[l], nums[min] = nums[min], nums[l]
			break
		}
		l--
	}

	for i, j := l+1, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

// -- end --
