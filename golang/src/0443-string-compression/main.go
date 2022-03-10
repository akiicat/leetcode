package main
import "strconv"

// T: O(n)
// M: O(1)
// -- start --

func compress(chars []byte) int {
  end, count := 0, 1

  for i, char := range chars {
    if i+1 < len(chars) && char == chars[i+1] {
      count++
    } else {
      chars[end], end = char, end + 1
      if count > 1 {
        for _, num := range strconv.Itoa(count) {
          chars[end], end = byte(num), end + 1
        }
      }
      count = 1
    }
  }

  return end
}

// -- end --

