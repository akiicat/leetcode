package main
import "testing"
import . "main/pkg/testing_helper"

func TestDivisorGame(t *testing.T) {
  i, o := [][]int{[]int{10,20},[]int{30,200},[]int{400,50},[]int{30,20}}, 110
  T(t, S(i), S(twoCitySchedCost(i)), S(o))

  i, o = [][]int{[]int{259,770},[]int{448,54},[]int{926,667},[]int{184,139},[]int{840,118},[]int{577,469}}, 1859
  T(t, S(i), S(twoCitySchedCost(i)), S(o))
}

func TestTwoCitySchedCostArray(t *testing.T) {
  i, o := [][]int{[]int{10,20},[]int{30,200},[]int{400,50},[]int{30,20}}, 110
  T(t, S(i), S(twoCitySchedCostArray(i)), S(o))

  i, o = [][]int{[]int{259,770},[]int{448,54},[]int{926,667},[]int{184,139},[]int{840,118},[]int{577,469}}, 1859
  T(t, S(i), S(twoCitySchedCostArray(i)), S(o))
}

