package main

// T: O(n + m) n is the length of commands, m is the length of obstacles
// M: O(m)
// -- start --

type point struct {
  x int
  y int
}

func robotSim(commands []int, obstacles [][]int) int {
  dx := []int{0, 1, 0, -1}
  dy := []int{1, 0, -1, 0}

  m := make(map[point]bool)

  for _, obs := range obstacles {
    m[point{obs[0], obs[1]}] = true
  }

  x, y, di := 0, 0, 0
  max := 0

  for _, step := range commands {
    if step == -1 {
      di = (di + 1) % 4
      continue
    }
    if step == -2 {
      di = (di + 3) % 4
      continue
    }

    for i := 0; i < step; i++ {
      nx := x + dx[di]
      ny := y + dy[di]

      if !m[point{nx, ny}] {
        x, y = nx, ny
        max = Max(max, x * x + y * y)
      }
    }
  }

  return max
}

func Max(a, b int) int {
  if a > b {
    return a
  }
  return b
}

// -- end --

