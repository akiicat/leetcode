package main

// T: O(n) n is the length of typed
// M: O(1)
// -- start --

func isLongPressedName(name string, typed string) bool {
  i := 0
  for j, _ := range typed {
    if name[i] != typed[j] {
      continue
    }

    i++
    if i == len(name) {
      return true
    }
  }

  return i == len(name)
}

// -- end --

