package main

// T: O(1) fix size
// M: O(1)
// -- start --

func numRookCaptures(board [][]byte) int {
  for x, row := range board {
    for y, v := range row {
      if v == 'R' {
        return Capture(board, x, y)
      }
    }
  }

  return 0
}

func Capture(board [][]byte, x, y int) int {

  dirs := [][]int{
    []int{0,1},
    []int{0,-1},
    []int{1,0},
    []int{-1,0},
  }

  sum := 0

  for _, d := range dirs {
    for r := 1; r < 8; r++ {
      dx := x + r * d[0]
      dy := y + r * d[1]

      if !(0 <= dx && dx < 8 && 0 <= dy && dy < 8) {
        break
      }

      if board[dx][dy] == '.' {
        continue
      }

      if board[dx][dy] == 'p' {
        sum++
        break
      }

      if board[dx][dy] == 'B' {
        break
      }
    }
  }

  return sum
}

// -- end --

