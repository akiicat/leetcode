package main

// T: O(n) n is the length of S
// M: O(1)
// -- start --

func numberOfLines(widths []int, S string) []int {
  line, index := 1, 0

  for _, v := range S {
    width := widths[v-'a']

    if index + width > 100 {
      line, index = line+1, 0 // reset
    }

    index += width
  }

  return []int{line, index}
}

// -- end --

