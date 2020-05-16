package main

// T: O(n)
// M: O(1)
// -- start --

var nums = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
var romans = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

func intToRoman(num int) string {
  s := ""

  for i, n := range nums {
    for num >= n {
      s += romans[i]
      num -= n
    }
  }

  return s
}


var m = map[int]string{
  0   : "",

  1   : "I",
  2   : "II",
  3   : "III",
  4   : "IV",
  5   : "V",
  6   : "VI",
  7   : "VII",
  8   : "VIII",
  9   : "IX",

  10  : "X",
  20  : "XX",
  30  : "XXX",
  40  : "XL",
  50  : "L",
  60  : "LX",
  70  : "LXX",
  80  : "LXXX",
  90  : "XC",

  100 : "C",
  200 : "CC",
  300 : "CCC",
  400 : "CD",
  500 : "D",
  600 : "DC",
  700 : "DCC",
  800 : "DCCC",
  900 : "CM",

  1000: "M",
  2000: "MM",
  3000: "MMM",
}

func intToRomanMap(num int) string {
  s := ""

  l := 1
  for num != 0 {
    s = m[l* (num % 10)] + s
    l, num = l * 10, num / 10
  }

  return s
}

// -- end --

