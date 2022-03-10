package main

// T: O(n)
// M: O(1)
// -- start --

func lengthOfLastWord(s string) int {
  start := len(s) - 1

  for start >= 0 {
    if s[start] != ' ' {
      break
    }
    start--
  }

  for i := start - 1; i >= 0; i-- {
    if s[i] == ' ' {
      return start - i
    }
  }

  return start + 1
}

func lengthOfLastWordBoolean(s string) int {
	length := 0
	foundCharacter := false

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if foundCharacter {
				break
			}

			continue
		}

		length++
		foundCharacter = true
	}

	return length
}

// -- end --

