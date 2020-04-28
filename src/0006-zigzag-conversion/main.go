package main

// T: O(n)
// M: O(1)
// -- start --

func convert(s string, numRows int) string {
  if numRows == 1 {
    return s
  }

  bs := []byte{}
  d := 2 * (numRows-1)

  for offset := 0; offset < numRows; offset++ {

    for i := offset; i < len(s); i += d {
      bs = append(bs, s[i])

      if 0 < offset && offset < numRows-1 {
        m := (numRows - offset - 1) * 2 // mirror
        if i + m < len(s) {
          bs = append(bs, s[i+m])
        }
      }
    }
  }

  return string(bs)
}

// -- end --

