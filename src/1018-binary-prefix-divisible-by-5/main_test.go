package main
import "testing"
import . "main/pkg/testing_helper"

func TestPrefixesDivBy5(t *testing.T) {
  i, o := []int{0,1,1}, []bool{true,false,false}
  T(t, S(i), S(prefixesDivBy5(i)), S(o))

  i, o = []int{1,1,1}, []bool{false,false,false}
  T(t, S(i), S(prefixesDivBy5(i)), S(o))

  i, o = []int{0,1,1,1,1,1}, []bool{true,false,false,false,true,false}
  T(t, S(i), S(prefixesDivBy5(i)), S(o))

  i, o = []int{1,1,1,0,1}, []bool{false,false,false,false,false}
  T(t, S(i), S(prefixesDivBy5(i)), S(o))

  i, o = []int{1,0,1,1,1,1,0,0,0,0,1,0,0,0,0,0,1,0,0,1,1,1,1,1,0,0,0,0,1,1,1,0,0,0,0,0,1,0,0,0,1,0,0,1,1,1,1,1,1,0,1,1,0,1,0,0,0,0,0,0,1,0,1,1,1,0,0,1,0}, []bool{false,false,true,false,false,false,false,false,false,false,true,true,true,true,true,true,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,true,false,false,false,true,false,false,true,false,false,true,true,true,true,true,true,true,false,false,true,false,false,false,false,true,true}
  T(t, S(i), S(prefixesDivBy5(i)), S(o))
}

