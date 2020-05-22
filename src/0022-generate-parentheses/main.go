package main

// T: O(4^n / n^(3/2))
// M: O(4^n / n^(3/2))
// -- start --

func generateParenthesis(n int) []string {
  if n == 1 {
    return []string{"()"}
  }

  res := []string{}
  R(&res, "", 0, 0, n)

  return res
}

func R(res *[]string, S string, l, r, n int) {
  if len(S) == 2 * n {
    *res = append(*res, S)
    return
  }

  if l < n {
    R(res, S+"(", l+1, r, n)
  }
  if r < l {
    R(res, S+")", l, r+1, n)
  }
}

// -- end --

