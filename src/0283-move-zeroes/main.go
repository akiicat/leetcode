package main

// T: O(n) n for the number of array
// M: O(1)
// -- start --

func moveZeroes(nums []int)  {
  cur := 0;

  for i, num := range nums {
    if num != 0 {
      nums[i], nums[cur] = nums[cur], nums[i]
      cur++
    }
  }
}

// -- end --

