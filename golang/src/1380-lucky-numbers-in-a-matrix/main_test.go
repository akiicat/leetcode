package main
import "testing"
import . "main/pkg/testing_helper"

func TestLuckyNumbers(_t *testing.T) {
  i, o := [][]int{
    []int{3,7,8},
    []int{9,11,13},
    []int{15,16,17},
  }, []int{15}
  T(_t, S(i), S(luckyNumbers(i)), S(o))

  i, o = [][]int{
    []int{1,10,4,2},
    []int{9,3,8,7},
    []int{15,16,17,12},
  }, []int{12}
  T(_t, S(i), S(luckyNumbers(i)), S(o))

  i, o = [][]int{
    []int{7,8},
    []int{1,2},
  }, []int{7}
  T(_t, S(i), S(luckyNumbers(i)), S(o))

  i, o = [][]int{
    []int{56216},
    []int{63251},
    []int{75772},
    []int{1945},
    []int{27014},
  }, []int{75772}
  T(_t, S(i), S(luckyNumbers(i)), S(o))
}

