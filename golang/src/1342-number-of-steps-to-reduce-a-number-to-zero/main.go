package main

// T: O(log(n))
// M: O(1)
// -- start --

func numberOfSteps(num int) int {
  c := 0
  for num != 0 {
    c++
    if num & 0x1 == 0 {
      num >>= 1
    } else {
      num--
    }
  }
  return c
}

// -- end --

