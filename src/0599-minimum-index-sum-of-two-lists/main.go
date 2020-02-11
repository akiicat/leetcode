package main

// T: O((n + m)*l)
// M: O(m*l)
// -- start --

// T: O((n + m)*l)
// M: O(m*l)
func findRestaurant(list1 []string, list2 []string) []string {
  if len(list1) > len(list2) {
    return findRestaurant(list2, list1)
  }

  m := make(map[string]int)
  for i, v := range list1 {
    m[v] = i
  }

  min := 1<<31
  rtn := []string{}
  for i, v := range list2 {
    c, ok := m[v]
    if !ok {
      continue
    }

    if c + i < min {
      min = c + i
      rtn = []string{}
    }

    if c + i == min {
      rtn = append(rtn, v)
    }
  }

  return rtn
}

// T: O((n + m)*l)
// M: O((n + m)*l)
func findRestaurantNegative(list1 []string, list2 []string) []string {
  m := make(map[string]int)

  for i, v := range list1 {
    m[v] = - i - 1
  }

  for i, v := range list2 {
    val, ok := m[v]
    if ok {
      m[v] = i - val
    }
  }

  min := 1<<31
  for _, v := range m {
    if v > 0 && v < min {
      min = v
    }
  }

  rtn := []string{}
  for k, v := range m {
    if v == min {
      rtn = append(rtn, k)
    }
  }

  return rtn
}

// -- end --

