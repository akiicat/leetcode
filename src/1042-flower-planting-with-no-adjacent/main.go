package main

// T: O(n)
// M: O(n)
// -- start --

func gardenNoAdj(N int, paths [][]int) []int {
  m := make([][]int, N)
  for _, v := range paths {
    x, y := v[0]-1, v[1]-1
    m[x] = append(m[x], y)
    m[y] = append(m[y], x)
  }

  rtn := make([]int, N)
  for i := 0; i < N; i++ {
    colors := [5]bool{}
    for _, j := range m[i] {
      colors[rtn[j]] = true
    }
    for c := 1; c <= 4; c++ {
      if colors[c] == false {
        rtn[i] = c
        break
      }
    }
  }

  return rtn
}

// -- end --

