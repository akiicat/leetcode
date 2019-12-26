package main
import "fmt"
import "sort"

func main() {
  i, o := []int{4,2,1,3}, [][]int{[]int{1,2}, []int{2,3}, []int{3,4}}
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %v\n", minimumAbsDifference(i))
  fmt.Printf("Expect: %v\n", o)

  i, o = []int{1,3,6,10,15}, [][]int{[]int{1,3}}
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %v\n", minimumAbsDifference(i))
  fmt.Printf("Expect: %v\n", o)

  i, o = []int{3,8,-10,23,19,-4,-14,27}, [][]int{[]int{-14,-10}, []int{19,23}, []int{23,27}}
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %v\n", minimumAbsDifference(i))
  fmt.Printf("Expect: %v\n", o)
}

// T: O(n*log(n))
// M: O(n)
// -- start --

func minimumAbsDifference(arr []int) [][]int {
  sort.Ints(arr)

  min := 1<<31
  rtn := [][]int{}
  for i := 1; i < len(arr); i++ {
    if arr[i] - arr[i-1] < min {
      min = arr[i] - arr[i-1]
      rtn = nil
    }
    if arr[i] - arr[i-1] == min {
      rtn = append(rtn, []int{arr[i-1], arr[i]})
    }
  }

  return rtn
}

func minimumAbsDifferenceTwoLoop(arr []int) [][]int {
  sort.Ints(arr)

  min := 1<<31
  for i := 1; i < len(arr); i++ {
    if arr[i] - arr[i-1] < min {
      min = arr[i] - arr[i-1]
    }
  }

  rtn := [][]int{}
  for i := 1; i < len(arr); i++ {
    if arr[i] - arr[i-1] == min {
      rtn = append(rtn, []int{arr[i-1], arr[i]})
    }
  }

  return rtn
}

// -- end --

