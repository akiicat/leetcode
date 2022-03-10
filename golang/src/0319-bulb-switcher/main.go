package main
import "math"

// T: O(log(n))
// M: O(1)
// -- start --

// http://potatomatooo.blogspot.com/2016/05/leetcode-319-bulb-switcher.html
func bulbSwitch(n int) int {
  return int(math.Sqrt(float64(n)))
}

// T: O(n ^ 2)
// M: O(n)
func bulbSwitchBruteForce(n int) int {
  m := make([]bool, n)

  sum := 0
  for i := 1; i <= n; i++ {
    for j := i-1; j < n; j += i {
      m[j] = !m[j]
    }
  }

  for i := 0; i < n; i++ {
    if m[i] {
      sum++
    }
  }

  return sum
}

// -- end --

