package main

// T: O(n * m)
// M: O(n + m)
// -- start --

func multiply(num1 string, num2 string) string {
  if num1[0] == '0' || num2[0] == '0' {
    return "0"
  }

  n1, n2 := []byte(num1), []byte(num2)
  rtn := make([]byte, len(num1) + len(num2))

  for i := len(num1)-1; i >= 0; i-- {
    for j := len(num2)-1; j >= 0; j--{
      rtn[i+j+1] += (n1[i]-'0') * (n2[j]-'0')

      if rtn[i+j+1] >= 0 {
        rtn[i+j] += rtn[i+j+1] / 10
        rtn[i+j+1] %= 10
      }
    }
  }

  if rtn[0] == 0 {
    rtn = rtn[1:]
  }

  for i := range rtn {
    rtn[i] += '0'
  }

  return string(rtn)
}

func multiplyReverse(num1 string, num2 string) string {
  n1, n2 := []byte(num1), []byte(num2)

  reverseString(n1)
  reverseString(n2)

  rtn := make([]byte, 222)

  for i := range n1 {
    n1[i] -= '0'
  }
  for i := range n2 {
    n2[i] -= '0'
  }

  for i := range n1 {
    for j := range n2 {
      rtn[i+j] += n1[i] * n2[j]
      rtn[i+j+1] += rtn[i+j] / 10
      rtn[i+j] %= 10
    }
  }

  for i := len(rtn)-1; i >= 0; i-- {
    if rtn[i] != 0 || i == 0 {
      rtn = rtn[:i+1]
      break;
    }
  }

  for i := range rtn {
    rtn[i] += '0'
  }

  reverseString(rtn)

  return string(rtn)
}

// leetcode 344. reverse string
func reverseString(s []byte) {
  n := len(s)
  for i := 0; i < n / 2; i++ {
    s[i], s[n-1-i] = s[n-1-i], s[i]
  }
}

// -- end --

