package main

// T: O(m*n)
// M: O(m*n)
// -- start --

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
  dirs := [][]int{
    []int{1,0},
    []int{-1,0},
    []int{0,1},
    []int{0,-1},
  }

  m := make([][]bool, R)
  for i := 0; i < R; i++ {
    m[i] = make([]bool, C)
  }

  m[r0][c0] = true
  q := [][]int{[]int{r0,c0}}

  for i := 0; i != len(q); i++ {
    for _, d := range dirs {
      r, c := q[i][0] + d[0], q[i][1] + d[1]
      if 0 <= r && r < R && 0 <= c && c < C && !m[r][c] {
        m[r][c] = true
        q = append(q, []int{r,c})
      }
    }
  }

  return q
}

// -- end --

