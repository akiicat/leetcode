package main
import "sort"

// T: O(n) n is len(S)
// M: O(1) for 26
// -- start --

// T: O(n)
// M: O(1)
type bucket struct {
	c   byte
	num int
}

type bucketSlice []bucket

func (bs bucketSlice) Len() int           { return len(bs) }
func (bs bucketSlice) Less(i, j int) bool { return bs[i].num < bs[j].num }
func (bs bucketSlice) Swap(i, j int)      { bs[i], bs[j] = bs[j], bs[i] }

func reorganizeString(S string) string {
	bs := make(bucketSlice, 26)
  for _, v := range S {
    i := v - 'a'
    if bs[i].num == 0 {
      bs[i].c = byte(v)
    }
		bs[i].num++
	}

  sort.Sort(sort.Reverse(bs))

  if bs[0].num > (len(S) + 1) / 2 {
    return ""
  }

  res := make([]byte, len(S))

  i := 0
  for _, b := range bs {
    for b.num > 0 {
      if i >= len(S) {
        i = 1
      }
      res[i] = b.c
      i, b.num = i+2, b.num-1
    }
  }

  return string(res)
}

// T: O(n)
// M: O(1)
func reorganizeStringFill(S string) string {
  m := [26]int{}
  for _, v := range S {
    m[v - 'a']++
  }

  rtn, cur := make([]byte, len(S)), 0
  for {
    mi := findMax(m)

    if mi == -1 {
      break
    }

    for cur < len(S) && m[mi] > 0 {
      rtn[cur] = byte('a' + mi)

      if cur & 0x1 == 1 && rtn[cur] == rtn[cur-1] {
        return ""
      }

      m[mi]--
      cur += 2

      if cur >= len(S) && cur & 0x1 == 0 {
        cur = 1
      }
    }
  }

  return string(rtn)
}

func findMax(m [26]int) int {
  mi := 0
  for i, v := range m {
    if v > m[mi] {
      mi = i
    }
  }

  if m[mi] == 0 {
    return -1
  }

  return mi
}

// T: O(n)
// M: O(1)
func reorganizeStringAppend(S string) string {
  m := [26]int{}

  for _, v := range S {
    m[v - 'a']++
  }

  s := ""
  for {

    m1, m2 := findTwoMax(m)

    if m1 == -1 {
      break
    }

    if m[m1] > 1 && m2 == -1 {
      return ""
    }

    if m[m1] == 1 && m2 == -1 {
      s += string([]byte{byte('a' + m1)})
      break
    }

    s += string([]byte{byte('a' + m1), byte('a' + m2)})
    m[m1]--
    m[m2]--
  }

  return s
}

func findTwoMax(m [26]int) (int, int) {
  m1, m2 := 0, 1
  for i, v := range m {
    if v > m[m1] {
      m1, m2 = i, m1
    } else if i != m1 && v > m[m2] {
      m2 = i
    }
  }
  if m[m1] == 0 {
    m1 = -1
  }

  if m[m2] == 0 {
    m2 = -1
  }

  return m1, m2
}

func Abs(a int) int {
  if a < 0 {
    return -a
  }
  return a
}

// -- end --

