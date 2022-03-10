package main

// T: O(n)
// M: O(1)
// -- start --

func plusOne(digits []int) []int {
  carry := 1
  for i := len(digits) - 1; i >= 0; i-- {
    digits[i] += carry

    // digits[i], carry = sum % 10, sum / 10
    if digits[i] >= 10 {
      digits[i] -= 10
    } else {
      carry = 0
    }

  }

  if carry == 1 {
    return append([]int{carry}, digits...)
  }

  return digits
}

// -- end --

