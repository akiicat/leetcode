package main

// T: O(n)
// M: O(1)
// -- start --

func largeGroupPositions(S string) [][]int {
  S = S + " "

  rtn := [][]int{}

  l := 0 // left
  for i, _ := range S {
    if S[l] == S[i] {
      continue
    }

    if i - l >= 3 {
      rtn = append(rtn, []int{l, i-1})
    }

    l = i
  }

  return rtn
}

// -- end --

