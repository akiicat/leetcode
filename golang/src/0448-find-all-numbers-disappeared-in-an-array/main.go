package main

// T: O(n)
// M: O(1)
// -- start --

// T: O(n)
// M: O(1)
func findDisappearedNumbers(nums []int) []int {
  rtn := make([]int, 0)

  for _, num := range nums {
    if num < 0 {
      num *= -1
    }
    if nums[num-1] > 0 {
      nums[num-1] *= -1
    }
  }

  for i, num := range nums {
    if num > 0 {
      rtn = append(rtn, i+1)
    }
  }

  return rtn
}

// T: O(n)
// M: O(n)
func findDisappearedNumbersMap(nums []int) []int {
  n := len(nums)
  m := make([]bool, n + 1)
  rtn := []int{}

  for _, num := range nums {
    m[num] = true
  }

  for i := 1; i < n+1; i++ {
    if m[i] == false {
      rtn = append(rtn, i)
    }
  }

  return rtn
}

// -- end --

