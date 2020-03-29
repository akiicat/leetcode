package main

// T: O(1)
// M: O(1)
// -- start --

func tictactoe(moves [][]int) string {
  m := [3][3]int{}

  for i, v := range moves {
    if i & 0x1 == 0 { // A
      m[v[0]][v[1]] = 1
    } else { // B
      m[v[0]][v[1]] = 2
    }
  }

  if is_win(m, 1) {
    return "A"
  }

  if is_win(m, 2) {
    return "B"
  }

  if len(moves) == 9 {
    return "Draw"
  }

  return "Pending"
}

func is_win(m [3][3]int, p int) bool {
  for i := 0; i < 3; i++ {
    if m[i][0] == p && m[i][1] == p && m[i][2] == p {
      return true
    }

    if m[0][i] == p && m[1][i] == p && m[2][i] == p {
      return true
    }
  }

  if m[0][0] == p && m[1][1] == p && m[2][2] == p {
    return true
  }

  if m[2][0] == p && m[1][1] == p && m[0][2] == p {
    return true
  }

  return false
}

// -- end --

