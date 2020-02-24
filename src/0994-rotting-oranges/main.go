package main

// T: O(m*n)
// M: O(m*n)
// -- start --

func orangesRotting(grid [][]int) int {
  m, n := len(grid), len(grid[0])
  q := []int{}
  dirs := [][]int{
    []int{0, 1},
    []int{1, 0},
    []int{0, -1},
    []int{-1, 0},
  }

  fresh := 0
  for x, row := range grid {
    for y, v := range row {
      if v == 2 {
        q = append(q, x, y)
      }
      if v != 0 {
        fresh++
      }
    }
  }

  if fresh == 0 {
    return 0
  }

  minute := 0
  for len(q) != 0 {
    t := []int{}

    for len(q) != 0 {
      fresh--
      for _, d := range dirs {
        rx := q[0] + d[0]
        ry := q[1] + d[1]
        if 0 <= rx && rx < m && 0 <= ry && ry < n && grid[rx][ry] == 1 {
          grid[rx][ry] = 2
          t = append(t, rx, ry)
        }
      }
      q = q[2:]
    }

    q = t
    minute++
  }

  if fresh != 0 {
    return -1
  }

  return minute-1
}

// T: O(m*n)
// M: O(m*n)
func orangesRottingBFS(grid [][]int) int {
  m, n := len(grid), len(grid[0])

  q := []int{}
  fresh := 0

  for x, row := range grid {
    for y, v := range row {
      if v == 2 {
        q = append(q, x, y)
      }
      if v != 0 {
        fresh++
      }
    }
  }

  if fresh == 0 {
    return 0
  }

  minute := 0
  for len(q) != 0 {
    t := q
    q = []int{}

    for len(t) != 0 {
      rx, ry := t[0], t[1]
      t = t[2:]
      fresh--

      if rx != 0 {
        q = append(q, Rotten(grid, rx-1, ry)...)
      }

      if ry != 0 {
        q = append(q, Rotten(grid, rx, ry-1)...)
      }

      if rx != m-1 {
        q = append(q, Rotten(grid, rx+1, ry)...)
      }

      if ry != n-1 {
        q = append(q, Rotten(grid, rx, ry+1)...)
      }
    }

    minute++
  }

  if fresh != 0 {
    return -1
  }

  return minute-1
}

func Rotten(grid [][]int, x, y int) []int {
  if grid[x][y] != 1 {
    return []int{}
  }
  grid[x][y] = 2
  return []int{x,y}
}

// -- end --

