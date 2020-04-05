package main

import "strconv"

// T: O(1)
// M: O(1)
// -- start --

func maximum69Number(num int) int {
  s := 0
  for ; num != 0; num /= 10 {
    s = 10 * s + num % 10
  }

  six := true
  for ; s != 0; s /= 10 {
    t := s % 10
    if t == 6 && six {
      t, six = 9, false
    }
    num = 10 * num + t
  }

  return num
}

func maximum69NumberStrconv(num int) int {
  s := strconv.Itoa(num)
  bs := []byte(s)
	for i, v := range bs {
		if v == '6' {
			bs[i] = '9'
			break
		}
	}
	res, _ := strconv.Atoi(string(bs))

	return res
}

// -- end --

