package main

// T: O(n)
// M: O(1)
// -- start --

func countLargestGroup(n int) int {
  m := [37]int{}

  for i := 1; i <= n; i++ {
    m[Sum(i)]++
  }

  max, c := 0, 0
  for _, v := range m {
    if v > max {
      max, c = v, 0
    }
    if v == max {
      c++
    }
  }

  return c
}

func Sum(a int) int {
  sum := 0
  for a != 0 {
    sum += a % 10
    a /= 10
  }
  return sum
}

// -- end --

