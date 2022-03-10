package main

// T: O(n)
// M: O(1)
// -- start --

// Two Pointer
// T: O(n)
// M: O(1)
func maxArea(height []int) int {
  max := 0

  l, r := 0, len(height)-1
  for l < r {
    max = Max(max, (r - l) * Min(height[l], height[r]))

    if height[l] < height[r] {
      l++
    } else {
      r--
    }
  }

  return max
}

// T: O(n ^ 2)
// M: O(n)
func maxAreaDynamicProgramming(height []int) int {
  m := make([]int, len(height))

  for d := 1; d < len(height); d++ {
    for i := d; i < len(height); i++ {
      m[i] = Max(m[i], m[i-1])
      m[i] = Max(m[i], d * Min(height[i], height[i-d]))
    }
  }

  return m[len(m)-1]
}

func Max(a, b int) int {
  if a > b {
    return a
  }
  return b
}

func Min(a, b int) int {
  if a < b {
    return a
  }
  return b
}

// -- end --

