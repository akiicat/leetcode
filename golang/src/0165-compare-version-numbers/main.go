package main

// T: O(n)
// M: O(1)
// -- start --

func compareVersion(version1 string, version2 string) int {
  i1, i2, v1, v2 := 0, 0, 0, 0
  for i1 < len(version1) && i2 < len(version2) {
    v1, i1 = getNextVersion(version1, i1)
    v2, i2 = getNextVersion(version2, i2)

    if v1 > v2 {
      return 1
    }
    if v1 < v2 {
      return -1
    }
  }

  for i1 < len(version1) {
    v1, i1 = getNextVersion(version1, i1)
    if v1 > 0 {
      return 1
    }
  }

  for i2 < len(version2) {
    v2, i2 = getNextVersion(version2, i2)
    if v2 > 0 {
      return -1
    }
  }

  return 0
}

func getNextVersion(v string, i int) (int, int) {
  r := 0
  for i < len(v) {
    if v[i] == '.' {
      i++
      break
    }

    r = 10 * r + int(v[i] - '0')
    i++
  }
  return r, i
}

// -- end --

