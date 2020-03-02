package main

// T: O(n)
// M: O(1)
// -- start --

func canThreePartsEqualSum(A []int) bool {
  sum := 0
  for _, v := range A {
    sum += v
  }
  if sum % 3 != 0 {
    return false
  }
  sum /= 3

  s, c := 0, 0
  for _, v := range A {
    s += v
    if s == (c+1)*sum {
      c++
    }
  }

  return c == 3
}

// -- end --

