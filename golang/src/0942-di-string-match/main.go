package main

// T: O(n)
// M: O(n)
// -- start --

func diStringMatch(S string) []int {
  n := len(S)
  min, max := 0, n
  rtn := make([]int, n+1, n+1)

  for i, v := range S {
    if v == 'I' {
      rtn[i] = min
      min++
    } else {
      rtn[i] = max
      max--
    }
  }

  rtn[n] = min

  return rtn
}

// -- end --

