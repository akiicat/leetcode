package main
import "sort"

// T: O((m + n) * log(n)) n is the number of heaters
// M: O(1)
// -- start --

// T: O((m + n) * log(n))
// M: O(1)
func findRadius(houses []int, heaters []int) int {
  sort.Ints(houses)
  sort.Ints(heaters)

  hi := 0
  max := 0

  for i := 0; i < len(houses); i++ {
    house := houses[i]

    for hi < len(heaters)-1 && Dist(house, heaters[hi]) >= Dist(house, heaters[hi+1]) {
      hi++
    }

    cur := Dist(house, heaters[hi])
    max = Max(max, cur)
  }

  return max
}

func Dist(a, b int) int {
  if a > b {
    return a - b
  }
  return b - a
}

// T: O(n * m)
// M: O(1)
func findRadiusBruteFroce(houses []int, heaters []int) int {
  max := 0

  for _, v := range houses {
    min := 1<<31-1
    for _, h := range heaters {
      min = Min(min, Abs(v - h))
    }
    max = Max(max, min)
  }

  return max
}

func Min(a, b int) int {
  if a > b {
    return b
  }
  return a
}

func Max(a, b int) int {
  if a > b {
    return a
  }
  return b
}

func Abs(a int) int {
  if a < 0 {
    return -a
  }
  return a
}

// -- end --

