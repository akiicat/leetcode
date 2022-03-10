package main

// T: O(n ** 2)
// M: O(1)
// -- start --

func flipAndInvertImage(A [][]int) [][]int {
  for _, row := range A {
    l, r := 0, len(row)-1
    for l <= r {
      row[l], row[r] = row[r] ^ 1, row[l] ^ 1

      l++
      r--
    }
  }
  return A
}

func flipAndInvertImageXOR(A [][]int) [][]int {
  for _, row := range A {
    l, r := 0, len(row)-1
    for l <= r {
      if l == r {
        row[l] ^= 1
      }

      row[l], row[r] = row[r], row[l]
      row[l] ^= 1
      row[r] ^= 1

      l++
      r--
    }
  }
  return A
}

// -- end --

