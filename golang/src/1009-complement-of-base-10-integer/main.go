package main

// T: O(1) for 32
// M: O(1)
// -- start --

func bitwiseComplement(N int) int {
  if N == 0 {
    return 1
  }

  mask := -1

  for mask & N != 0 {
    mask <<= 1
  }

  return N ^ (^mask)
}

// -- end --

