package main

// T: O(n)
// M: O(1)
// -- start --

func reverse(x int) int {
  input := x // speed up, do not use x as computation
  if input != int(int32(input)) {
    return 0
  }

  reverse := 0
  for input != 0 {
    reverse = 10 * reverse + input % 10
    input = input / 10
  }

  if reverse != int(int32(reverse)) {
    return 0
  }

  return reverse
}

// -- end --

