package main

// T: O(n)
// M: O(1)
// -- start --

func findNumbers(nums []int) int {
  sum := 0
  for _, v := range nums {
    if is_even(v) {
      sum++
    }
  }
  return sum
}

func is_even(v int) bool {
  sum := 0

  for v != 0 {
    sum++
    v /= 10
  }

  return sum & 0x1 == 0
}

// -- end --

