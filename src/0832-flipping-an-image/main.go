package main
import "fmt"

func main() {
  i, o := [][]int{
    []int{1,1,0},
    []int{1,0,1},
    []int{0,0,0},
  },
  [][]int{
    []int{1,0,0},
    []int{0,1,0},
    []int{1,1,1},
  }
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %v\n", flipAndInvertImage(i))
  fmt.Printf("Expect: %v\n", o)

  i, o = [][]int{
    []int{1,1,0,0},
    []int{1,0,0,1},
    []int{0,1,1,1},
    []int{1,0,1,0},
  },
  [][]int{
    []int{1,1,0,0},
    []int{0,1,1,0},
    []int{0,0,0,1},
    []int{1,0,1,0},
  }
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %v\n", flipAndInvertImage(i))
  fmt.Printf("Expect: %v\n", o)
}

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

