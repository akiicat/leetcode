package main
import "testing"
import . "main/pkg/testing_helper"

func TestNextGreatestLetter(_t *testing.T) {
  i, t, o := []byte{'c','f','j'}, byte('a'), byte('c')
  T(_t, S(i, t), S(nextGreatestLetter(i, t)), S(o))

  i, t, o = []byte{'c','f','j'}, byte('c'), byte('f')
  T(_t, S(i, t), S(nextGreatestLetter(i, t)), S(o))

  i, t, o = []byte{'c','f','j'}, byte('d'), byte('f')
  T(_t, S(i, t), S(nextGreatestLetter(i, t)), S(o))

  i, t, o = []byte{'c','f','j'}, byte('g'), byte('j')
  T(_t, S(i, t), S(nextGreatestLetter(i, t)), S(o))

  i, t, o = []byte{'c','f','j'}, byte('j'), byte('c')
  T(_t, S(i, t), S(nextGreatestLetter(i, t)), S(o))

  i, t, o = []byte{'c','f','j'}, byte('k'), byte('c')
  T(_t, S(i, t), S(nextGreatestLetter(i, t)), S(o))
}

func TestNextGreatestLetterLinear(_t *testing.T) {
  i, t, o := []byte{'c','f','j'}, byte('a'), byte('c')
  T(_t, S(i, t), S(nextGreatestLetterLinear(i, t)), S(o))

  i, t, o = []byte{'c','f','j'}, byte('c'), byte('f')
  T(_t, S(i, t), S(nextGreatestLetterLinear(i, t)), S(o))

  i, t, o = []byte{'c','f','j'}, byte('d'), byte('f')
  T(_t, S(i, t), S(nextGreatestLetterLinear(i, t)), S(o))

  i, t, o = []byte{'c','f','j'}, byte('g'), byte('j')
  T(_t, S(i, t), S(nextGreatestLetterLinear(i, t)), S(o))

  i, t, o = []byte{'c','f','j'}, byte('j'), byte('c')
  T(_t, S(i, t), S(nextGreatestLetterLinear(i, t)), S(o))

  i, t, o = []byte{'c','f','j'}, byte('k'), byte('c')
  T(_t, S(i, t), S(nextGreatestLetterLinear(i, t)), S(o))
}

