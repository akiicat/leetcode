package main
import "testing"
import . "main/pkg/testing_helper"

func TestFindDisappearedNumbers(_t *testing.T) {
  i, o := []int{4,3,2,7,8,2,3,1}, []int{5,6}
  T(_t, S(i), S(findDisappearedNumbers(i)), S(o))

  i, o = []int{4,7,2,7,8,8,3,1}, []int{5,6}
  T(_t, S(i), S(findDisappearedNumbers(i)), S(o))
}

func TestFindDisappearedNumbersMap(_t *testing.T) {
  i, o := []int{4,3,2,7,8,2,3,1}, []int{5,6}
  T(_t, S(i), S(findDisappearedNumbersMap(i)), S(o))

  i, o = []int{4,7,2,7,8,8,3,1}, []int{5,6}
  T(_t, S(i), S(findDisappearedNumbersMap(i)), S(o))
}
