package main

// T: O(n)
// M: O(1)
// -- start --

func replaceElements(arr []int) []int {
  max := -1
  for i := len(arr)-1; i >=0 ; i-- {
    if arr[i] > max {
      max, arr[i] = arr[i], max
    } else {
      arr[i] = max
    }
  }
  return arr
}

func replaceElementsLeftToRight(arr []int) []int {
  max := arr[0]
  for i, v := range arr {
    if v == max {
      max = Max(arr, i+1)
    }
    arr[i] = max
  }
  return arr
}

func Max(arr []int, i int) int {
  max := -1
  for i < len(arr) {
    if arr[i] > max {
      max = arr[i]
    }
    i++
  }
  return max
}

// -- end --

