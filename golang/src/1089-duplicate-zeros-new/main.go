package main

// T: O(n)
// M: O(1)
// -- start --

func duplicateZeros(arr []int)  {
  i, c := 0, 0
  for i < len(arr) {
    if arr[c] == 0 {
      i++
    }
    i, c = i+1, c+1
  }
  i, c = i-1, c-1

  if i == len(arr) { // half zero
    arr[i-1] = 0
    i, c = i-2, c-1
  }

  for i > 0 {
    if arr[c] == 0 {
      arr[i] = 0
      i--
    }

    arr[i] = arr[c]
    i, c = i-1, c-1
  }
}

// T: O(n * n)
// M: O(1)
func duplicateZerosCopy(arr []int)  {
  for i:= 0; i < len(arr); i++ {
    if arr[i] == 0 {
      copy(arr[i+1:], arr[i:])
      i++
    }
  }
}

// -- end --
