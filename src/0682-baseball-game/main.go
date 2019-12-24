package main
import "fmt"
import "strconv"

func main() {
  i, o := []string{"5","2","C","D","+"}, 30
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %v\n", calPoints(i))
  fmt.Printf("Expect: %v\n", o)

  i, o = []string{"5","-2","4","C","D","9","+","+"}, 27
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %v\n", calPoints(i))
  fmt.Printf("Expect: %v\n", o)
}

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

