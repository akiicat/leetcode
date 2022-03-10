package main

// T: O(log(n)) n for the number of bound
// M: O(1)
// -- start --

func powerfulIntegers(x int, y int, bound int) []int {
  m := make(map[int]bool)
  for rx := 1; rx <= bound; rx = rx * x {
    for ry := 1; rx + ry <= bound; ry = ry * y {
      m[rx + ry] = true
      if y == 1 {
        break
      }
    }
    if x == 1 {
      break
    }
  }

  rtn := []int{}
  for k, _ := range m {
    rtn = append(rtn, k)
  }

  return rtn
}

// Power
// T: O(log(n)) n for the number of bound
// M: O(1)
func powerfulIntegersPower(x int, y int, bound int) []int {
  m := make(map[int]bool)

  i, rx := 1, 1
  for rx <= bound {

    j, ry := 1, 1
    for rx + ry <= bound {
      m[rx + ry] = true

      if y == 1 && j == 1 {
        break
      }

      j, ry = j+1, Pow(y, j)
    }

    if x == 1 && i == 1 {
      break
    }
    i, rx = i+1, Pow(x, i)
  }

  rtn := []int{}
  for k, _ := range m {
    rtn = append(rtn, k)
  }

  return rtn
}

// leetcode 50.
func Pow(x int, n int) int {
  if n == 0 {
    return 1
  }

  res, base := 1, x
  for n > 1 {
    if n & 0x1 == 0 {
      n, base = n >> 1, base * base
    } else {
      n, res = n - 1, res * base
    }
  }

  return res * base
}

// -- end --

