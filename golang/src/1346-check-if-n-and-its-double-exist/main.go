package main

// T: O(n)
// M: O(n)
// -- start --

func checkIfExist(arr []int) bool {
  m := map[int]bool{}

  for _, v := range arr {
    if m[v<<1] || ((v & 0x1 == 0) && m[v>>1]) {
      return true
    }
    m[v] = true
  }
  return false
}

// -- end --

