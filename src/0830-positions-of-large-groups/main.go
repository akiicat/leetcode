package main
import "fmt"

func main() {
  i, o := "abbxxxxzyy", [][]int{[]int{3,6}}
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %v\n", largeGroupPositions(i))
  fmt.Printf("Expect: %v\n", o)

  i, o = "abc", [][]int{}
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %v\n", largeGroupPositions(i))
  fmt.Printf("Expect: %v\n", o)

  i, o = "abcdddeeeeaabbbcd", [][]int{[]int{3,5},[]int{6,9},[]int{12,14}}
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %v\n", largeGroupPositions(i))
  fmt.Printf("Expect: %v\n", o)

  i, o = "aaa", [][]int{[]int{0,2}}
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %v\n", largeGroupPositions(i))
  fmt.Printf("Expect: %v\n", o)
}

// T: O(n)
// M: O(1)
// -- start --

func largeGroupPositions(S string) [][]int {
  S = S + " "

  rtn := [][]int{}

  l := 0 // left
  for i, _ := range S {
    if S[l] == S[i] {
      continue
    }

    if i - l >= 3 {
      rtn = append(rtn, []int{l, i-1})
    }

    l = i
  }

  return rtn
}

// -- end --

