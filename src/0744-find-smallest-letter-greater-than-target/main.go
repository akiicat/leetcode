package main
import "fmt"

func main() {
  i, t, o := []byte{'c','f','j'}, byte('a'), byte('c')
  fmt.Printf("Input:  %s\n", string(i))
  fmt.Printf("Output: %s\n", string(nextGreatestLetter(i, t)))
  fmt.Printf("Expect: %s\n", string(o))

  i, t, o = []byte{'c','f','j'}, byte('c'), byte('f')
  fmt.Printf("Input:  %s\n", string(i))
  fmt.Printf("Output: %s\n", string(nextGreatestLetter(i, t)))
  fmt.Printf("Expect: %s\n", string(o))

  i, t, o = []byte{'c','f','j'}, byte('d'), byte('f')
  fmt.Printf("Input:  %s\n", string(i))
  fmt.Printf("Output: %s\n", string(nextGreatestLetter(i, t)))
  fmt.Printf("Expect: %s\n", string(o))

  i, t, o = []byte{'c','f','j'}, byte('g'), byte('j')
  fmt.Printf("Input:  %s\n", string(i))
  fmt.Printf("Output: %s\n", string(nextGreatestLetter(i, t)))
  fmt.Printf("Expect: %s\n", string(o))

  i, t, o = []byte{'c','f','j'}, byte('j'), byte('c')
  fmt.Printf("Input:  %s\n", string(i))
  fmt.Printf("Output: %s\n", string(nextGreatestLetter(i, t)))
  fmt.Printf("Expect: %s\n", string(o))

  i, t, o = []byte{'c','f','j'}, byte('k'), byte('c')
  fmt.Printf("Input:  %s\n", string(i))
  fmt.Printf("Output: %s\n", string(nextGreatestLetter(i, t)))
  fmt.Printf("Expect: %s\n", string(o))
}

// T: O(log(n))
// M: O(1)
// -- start --

// T: O(log(n))
// M: O(1)
func nextGreatestLetter(letters []byte, target byte) byte {
  l, h := 0, len(letters)

  for l < h {
    m := (l + h) / 2
    if letters[m] <= target {
      l = m+1
    } else {
      h = m
    }
  }

  return letters[l % len(letters)]
}

// T: O(log(n))
// M: O(1)
func nextGreatestLetterLinear(letters []byte, target byte) byte {
  for _, v := range letters {
    if v > target {
      return byte(v)
    }
  }

  return letters[0]
}

// -- end --

