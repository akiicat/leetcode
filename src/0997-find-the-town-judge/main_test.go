package main
import "testing"
import . "main/pkg/testing_helper"

func TestFindJudge(_t *testing.T) {
  n, i, o := 2, [][]int{[]int{1,2}}, 2
  T(_t, S(n, i), S(findJudge(n, i)), S(o))

  n, i, o = 2, [][]int{[]int{2,1}}, 1
  T(_t, S(n, i), S(findJudge(n, i)), S(o))

  n, i, o = 3, [][]int{[]int{1,3},[]int{2,3}}, 3
  T(_t, S(n, i), S(findJudge(n, i)), S(o))

  n, i, o = 3, [][]int{[]int{1,3},[]int{2,3},[]int{3,1}}, -1
  T(_t, S(n, i), S(findJudge(n, i)), S(o))

  n, i, o = 3, [][]int{[]int{1,2},[]int{2,3}}, -1
  T(_t, S(n, i), S(findJudge(n, i)), S(o))

  n, i, o = 4, [][]int{[]int{1,3},[]int{1,4},[]int{2,3},[]int{2,4},[]int{4,3}}, 3
  T(_t, S(n, i), S(findJudge(n, i)), S(o))
}

func TestFindJudgeTwoLoop(_t *testing.T) {
  n, i, o := 2, [][]int{[]int{1,2}}, 2
  T(_t, S(n, i), S(findJudgeTwoLoop(n, i)), S(o))

  n, i, o = 2, [][]int{[]int{2,1}}, 1
  T(_t, S(n, i), S(findJudgeTwoLoop(n, i)), S(o))

  n, i, o = 3, [][]int{[]int{1,3},[]int{2,3}}, 3
  T(_t, S(n, i), S(findJudgeTwoLoop(n, i)), S(o))

  n, i, o = 3, [][]int{[]int{1,3},[]int{2,3},[]int{3,1}}, -1
  T(_t, S(n, i), S(findJudgeTwoLoop(n, i)), S(o))

  n, i, o = 3, [][]int{[]int{1,2},[]int{2,3}}, -1
  T(_t, S(n, i), S(findJudgeTwoLoop(n, i)), S(o))

  n, i, o = 4, [][]int{[]int{1,3},[]int{1,4},[]int{2,3},[]int{2,4},[]int{4,3}}, 3
  T(_t, S(n, i), S(findJudgeTwoLoop(n, i)), S(o))
}

