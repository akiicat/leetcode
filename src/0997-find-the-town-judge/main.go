package main

// T: O(n)
// M: O(n)
// -- start --

// T: O(n)
// M: O(n)
func findJudge(N int, trust [][]int) int {
  m := make([]int, N)

  r := 0

  for _, v := range trust {
    a, b := v[0]-1, v[1]-1
    m[a]--
    m[b]++

    if m[b] == N-1 {
      r = b
    }
  }

  if m[r] == N-1 {
    return r+1
  }

  return -1
}

// T: O(n + t)
// M: O(n)
func findJudgeTwoLoop(N int, trust [][]int) int {
  m := make([]int, N)

  for _, v := range trust {
    m[v[0]-1]--
    m[v[1]-1]++
  }

  for i := 0; i < N; i++ {
    if m[i] == N-1 {
      return i+1
    }
  }

  return -1
}

// -- end --

