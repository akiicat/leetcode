package main
import "fmt"

func main() {
  i, o := "IDID", []int{0,4,1,3,2}
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %v\n", diStringMatch(i))
  fmt.Printf("Expect: %v\n", o)

  i, o = "III", []int{0,1,2,3}
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %v\n", diStringMatch(i))
  fmt.Printf("Expect: %v\n", o)

  i, o = "DDI", []int{3,2,0,1}
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %v\n", diStringMatch(i))
  fmt.Printf("Expect: %v\n", o)
}

// T: O(n)
// M: O(n)
// -- start --

func diStringMatch(S string) []int {
  n := len(S)
  min, max := 0, n
  rtn := make([]int, n+1, n+1)

  for i, v := range S {
    if v == 'I' {
      rtn[i] = min
      min++
    } else {
      rtn[i] = max
      max--
    }
  }

  rtn[n] = min

  return rtn
}

// -- end --

