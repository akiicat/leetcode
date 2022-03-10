package main
import "strconv"

// T: O(n)
// M: O(n)
// -- start --

func calPoints(ops []string) int {
  s := []int{}
  for _, v := range ops {
    switch v {
    case "C":
      s = s[:len(s)-1]
    case "D":
      s = append(s, s[len(s)-1] * 2)
    case "+":
      s = append(s, s[len(s)-1] + s[len(s)-2])
    default:
      i, _ := strconv.Atoi(v)
      s = append(s, i)
    }
  }

  sum := 0
  for _, v := range s {
    sum += v
  }

  return sum
}

// -- end --

