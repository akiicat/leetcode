package main

// T: O(n)
// M: O(1)
// -- start --

// BFS
// T: O(n)
// M: O(1)
func shortestToChar(S string, C byte) []int {
  m := make(map[int]bool)

  q := []int{}
  for i := 0; i < len(S); i++ {
    if S[i] == C {
      q = append(q, i)
      m[i] = true
    }
  }

  rtn := make([]int, len(S))
  level := 0
  for len(q) != 0 {
    n := len(q)

    for _, i := range q {
      rtn[i] = level
      if i-1 >= 0 && !m[i-1] {
        q = append(q, i-1)
        m[i-1] = true
      }
      if i+1 < len(S) && !m[i+1] {
        q = append(q, i+1)
        m[i+1] = true
      }
    }

    q = q[n:]
    level++
  }

  return rtn
}

// T: O(n)
// M: O(1)
func shortestToCharTwoPass(S string, C byte) []int {
  rtn := make([]int, len(S))

  l := -1<<31
  for i := 0; i < len(S); i++ {
    if S[i] == C {
      l = i
    }
    rtn[i] = i - l
  }

  r := 1<<31
  for i := len(S)-1; i >= 0; i-- {
    if S[i] == C {
      r = i
    }
    rtn[i] = Min(rtn[i], r - i)
  }

  return rtn
}

func Min(a, b int) int {
  if a > b {
    return b
  }
  return a
}

// -- end --

