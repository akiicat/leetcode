package main

// T: O(n)
// M: O(1)
// -- start --

// T: O(n)
// M: O(1)
func maxSubArray(nums []int) int {
  m, cur := nums[0], 0

  for _, v := range nums {
    cur = max(cur+v, v)
    m = max(cur, m)
  }

  return m // max
}

func max(a, b int) int {
  if a > b {
    return a
  }
  return b
}

// T: O(n**2)
// M: O(n)
func maxSubArrayDynamicProgramming(nums []int) int {

  m := make([]int, len(nums))

  max := -1<<31

  for i := 0; i < len(nums); i++ {
    m[i] = 0
  }

  for d := 0; d < len(nums); d++ { // interval
    for i := 0; i < len(nums) - d; i++ {
      m[i] = m[i] + nums[i+d]
      if m[i] > max {
        max = m[i]
      }
    }
  }

  return max
}

// T: O(n ^ 2)
// T: O(1)
func maxSubArrayBruteForce(nums []int) int {

  max := -1<<31

  for i := 0; i < len(nums); i++ {
    for j := 0; j < len(nums) - i; j++ {
      tmp := sum(nums, j, j+i+1)

      if tmp > max {
        max = tmp
      }
    }
  }

  return max
}

func sum(nums []int, l, r int) int {
  s := 0
  for i := l; i < r; i++ {
    s += nums[i]
  }
  return s
}

// -- end --

