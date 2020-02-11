package main
import "sort"

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

