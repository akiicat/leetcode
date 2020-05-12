package nested_integer
import "strconv"

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
  val int
  list []*NestedInteger
  set bool
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (n NestedInteger) IsInteger() bool {
  return n.set
}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (n NestedInteger) GetInteger() int {
  return n.val
}

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {
  n.val = value
  n.set = true
}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (n *NestedInteger) Add(elem NestedInteger) {
  n.list = append(n.list, &elem)
}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (n NestedInteger) GetList() []*NestedInteger {
  return n.list
}

func (l *NestedInteger) ToStr() string {
  if l == nil {
    return ""
  }

  if l.set {
    return strconv.Itoa(l.val)
  }

  s := "["
  for i, v := range l.list {
    s += v.ToStr()
    if i + 1 != len(l.list) { // not last
      s += ","
    }
  }
  s += "]"

  return s
}

func NewNestedInteger(s string) *NestedInteger {
  return deserialize(s)
}

// leetcode 385.
func deserialize(s string) *NestedInteger {
  res := &NestedInteger{}

  if s[0] != '[' {
    v, _ := strconv.Atoi(s)
    res.SetInteger(v)
    return res
  }

  s = s[1:len(s)-1]

  last, count := 0, 0

  for i := 0; i <= len(s); i++ {
    if (i == len(s) || s[i] == ',') && count == 0 {
      res.Add(*(deserialize(s[last:i])))
      last = i+1
    } else if s[i] == '[' {
      count++
    } else if s[i] == ']' {
      count--
    }
  }

  return res
}


